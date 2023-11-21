package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type VideoInfo struct {
	GroupTitle string `json:"groupTitle"`
	Title      string `json:"title"`
}

func removeLeadingZeros(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	leadingZeros := data[:9]
	data = data[9:] // remove the first 9 bytes (zeros)

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return nil, err
	}

	return leadingZeros, nil
}

func addLeadingZeros(filename string, leadingZeros []byte) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data = append(leadingZeros, data...)

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func sanitizeName(name string) string {
	name = strings.ReplaceAll(name, "/", "-")
	name = strings.ReplaceAll(name, " ", "-")
	return name
}

func processDirectory(inputDir, outputDir string, keepOriginal, dryRun bool) {
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var m4sFiles []string
	var leadingZeros [][]byte

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".m4s") {
			m4sFiles = append(m4sFiles, file.Name())
		}
	}

	if len(m4sFiles) == 0 {
		fmt.Println("Could not find any .m4s files in the directory")
		return
	}

	videoInfoFile := inputDir + "/.videoInfo"
	videoInfoData, err := ioutil.ReadFile(videoInfoFile)
	if err != nil {
		fmt.Println("Error reading video info file:", err)
		return
	}

	var videoInfo VideoInfo
	err = json.Unmarshal(videoInfoData, &videoInfo)
	if err != nil {
		fmt.Println("Error parsing video info file:", err)
		return
	}

	videoInfo.GroupTitle = sanitizeName(videoInfo.GroupTitle)
	videoInfo.Title = sanitizeName(videoInfo.Title)

	outputDir = filepath.Join(outputDir, videoInfo.GroupTitle)
	if !dryRun {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			return
		}
	}

	outputFile := filepath.Join(outputDir, videoInfo.Title+".mp4")

	if _, err := os.Stat(outputFile); err == nil {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Output file already exists. Do you want to keep the original files? (yes/no): ")
		response, _ := reader.ReadString('\n')
		response = strings.ToLower(strings.TrimSpace(response))

		if response != "yes" {
			fmt.Println("Removing original files...")
			os.RemoveAll(inputDir)
			return
		} else {
			fmt.Println("Happy Hacking!")
			return
		}

	} else if !os.IsNotExist(err) {
		fmt.Println("Error checking if output file exists:", err)
		return
	}

	for i, file := range m4sFiles {
		m4sFiles[i] = inputDir + "/" + file

		if !dryRun {
			zeros, err := removeLeadingZeros(m4sFiles[i])
			if err != nil {
				fmt.Println("Error removing leading zeros from file:", err)
				return
			}
			leadingZeros = append(leadingZeros, zeros)
		}
	}

	args := make([]string, 0, 2*len(m4sFiles)+3)
	for _, file := range m4sFiles {
		args = append(args, "-i", file)
	}
	args = append(args, "-c", "copy", outputFile)

	if dryRun {
		fmt.Println("Would run: ffmpeg", strings.Join(args, " "))
	} else {
		cmd := exec.Command("ffmpeg", args...)
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error running ffmpeg:", err)
			return
		}

		fmt.Println("Successfully created output file:", outputFile)
	}

	if keepOriginal && !dryRun {
		for i, file := range m4sFiles {
			err = addLeadingZeros(file, leadingZeros[i])
			if err != nil {
				fmt.Println("Error adding leading zeros back to file:", err)
				return
			}
		}
	} else if !keepOriginal && !dryRun {
		os.RemoveAll(inputDir)
	}
}

func main() {
	inputDir := flag.String("inputDir", os.Getenv("HOME")+"/Movies/bilibili", "Input directory")
	outputDir := flag.String("outputDir", os.Getenv("HOME")+"/Tmp/Movies", "Output directory")
	convert := flag.String("convert", "no", "Whether to run the conversion or not")
	dryRun := flag.Bool("dry-run", false, "Whether to perform a dry run")

	flag.Parse()

	if strings.ToLower(*convert) != "yes" {
		fmt.Println("Usage: go run main.go -convert=yes [-inputDir=/path/to/input] [-outputDir=/path/to/output] [--dry-run]")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to keep the original directories? (yes/no): ")
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	keepOriginal := response == "yes"

	dirs, err := ioutil.ReadDir(*inputDir)
	if err != nil {
		fmt.Println("Error reading input directory:", err)
		return
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			processDirectory(filepath.Join(*inputDir, dir.Name()), *outputDir, keepOriginal, *dryRun)
		}
	}
}
