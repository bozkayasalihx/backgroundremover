package main

import (
    "os/exec"
)

func main() {
    inputVideo := "input.mp4"
    outputVideo := "output.mp4"
    greenScreenColor := "0x00FF00"

    cmd := exec.Command("ffmpeg", "-i", inputVideo, "-filter_complex",
        "chromakey=0.1:0.8:"+greenScreenColor, outputVideo)

    if err := cmd.Run(); err != nil {
        panic(err)
    }
}

