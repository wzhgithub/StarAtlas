package main

/*
 #include <stdio.h>
 #include <unistd.h>
 #include <termios.h>
 char getch(){
	 char ch = 0;
	 struct termios old = {0};
	 fflush(stdout);
	 if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
	 old.c_lflag &= ~ICANON;
	 old.c_lflag &= ~ECHO;
	 old.c_cc[VMIN] = 1;
	 old.c_cc[VTIME] = 0;
	 if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
	 if( read(0, &ch,1) < 0 ) perror("read()");
	 old.c_lflag |= ICANON;
	 old.c_lflag |= ECHO;
	 if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
	 return ch;
 }
*/
import "C"

// stackoverflow.com/questions/14094190/golang-function-similar-to-getchar

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/nl8590687/asrt-sdk-go/common"
	"github.com/nl8590687/asrt-sdk-go/sdk"
	wave "github.com/zenwerk/go-wave"
)

func main() {
	recordAudio()
	// httpDemo()
	// grpcDemo()
}

func recordAudio() {
	audioFileName := "testData/record.wav"
	fmt.Println("Recording. Press ESC to quit.")

	waveFile, err := os.Create(audioFileName)
	chk(err)

	// www.people.csail.mit.edu/hubert/pyaudio/  - under the Record tab
	inputChannels := 1
	outputChannels := 0
	sampleRate := 16000
	framesPerBuffer := make([]int16, 64)

	// init PortAudio
	portaudio.Initialize()
	//defer portaudio.Terminate()
	stream, err := portaudio.OpenDefaultStream(inputChannels, outputChannels, float64(sampleRate), len(framesPerBuffer), framesPerBuffer)
	chk(err)
	//defer stream.Close()

	// setup Wave file writer
	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputChannels,
		SampleRate:    sampleRate,
		BitsPerSample: 16, // if 16, change to WriteSample16()
	}

	waveWriter, err := wave.NewWriter(param)
	chk(err)

	//defer waveWriter.Close()
	go func() {
		key := C.getch()
		fmt.Println()
		fmt.Println("Cleaning up ...")
		if key == 27 {
			// better to control
			// how we close then relying on defer
			waveWriter.Close()
			stream.Close()
			portaudio.Terminate()
			fmt.Println("Play", audioFileName, "with a audio player to hear the result.")
			httpDemo()
			os.Exit(0)
		}

	}()

	// recording in progress ticker. From good old DOS days.
	ticker := []string{
		"-",
		"\\",
		"/",
		"|",
	}
	rand.Seed(time.Now().UnixNano())

	// start reading from microphone
	chk(stream.Start())
	for {
		chk(stream.Read())

		fmt.Printf("\rRecording is live now. Say something to your microphone! [%v]", ticker[rand.Intn(len(ticker)-1)])

		// write to wave file
		_, err := waveWriter.WriteSample16([]int16(framesPerBuffer)) // WriteSample16 for 16 bits
		chk(err)
	}
	chk(stream.Stop())
}

func httpDemo() {
	// 初始化
	host := "127.0.0.1"
	port := "20001"
	protocol := "http"

	sr := sdk.GetSpeechRecognizer(host, port, protocol)
	// ======================================================
	// 识别文件
	filename := "testData/record.wav"
	resultFile, err := sr.RecogniteFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	for index, res := range resultFile {
		fmt.Println("Wav文件语音识别结果 ", index, ":", res.Result)
	}

	byteData := sdk.LoadFile(filename)
	wave, err := sdk.DecodeWav(byteData)
	if err != nil {
		fmt.Println(err)
	}
	// ======================================================
	// 识别一段Wave音频序列
	result, err := sr.Recognite(wave.GetRawSamples(), wave.FrameRate, wave.Channels, wave.SampleWidth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语音识别结果：", result.Result)
	// ======================================================
	// 调用声学模型识别一段Wave音频序列
	result, err = sr.RecogniteSpeech(wave.GetRawSamples(), wave.FrameRate, wave.Channels, wave.SampleWidth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语音识别声学模型结果：", result.Result)
	// // ======================================================
	// // 调用语言模型1
	// pinyinResult := []string{}
	// for i := 0; i < len(result.Result.([]interface{})); i += 1 {
	// 	pinyinResult = append(pinyinResult, result.Result.([]interface{})[i].(string))
	// }

	// result, err = sr.RecogniteLanguage(pinyinResult)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("语言模型结果：", result.Result)
	// ======================================================
	// 调用语言模型2
	// sequencePinyin := []string{"ni3", "hao3", "a1"}
	// result, err = sr.RecogniteLanguage(sequencePinyin)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("语言模型结果：", result.Result)
}

func grpcDemo() {
	// 初始化
	host := "127.0.0.1"
	port := "20002"
	protocol := "grpc"

	sr := sdk.GetSpeechRecognizer(host, port, protocol)
	fmt.Println("sr:", sr)
	// ======================================================
	// 识别文件
	filename := "testData/data1.wav"
	resultFile, err := sr.RecogniteFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	for index, res := range resultFile {
		fmt.Println("Wav文件语音识别结果 ", index, ":", res.Result)
	}

	byteData := sdk.LoadFile(filename)
	wave, err := sdk.DecodeWav(byteData)
	if err != nil {
		fmt.Println(err)
	}
	// ======================================================
	// 识别一段Wave音频序列
	result, err := sr.Recognite(wave.GetRawSamples(), wave.FrameRate, wave.Channels, wave.SampleWidth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语音识别结果：", result.Result)
	// ======================================================
	// 识别一段长Wave音频序列
	longSample := wave.GetRawSamples()
	longSample = append(longSample, wave.GetRawSamples()...)
	resultLong, err := sr.RecogniteLong(longSample, wave.FrameRate, wave.Channels, wave.SampleWidth)
	if err != nil {
		fmt.Println(err)
	}

	for index, res := range resultLong {
		fmt.Println("长文件语音识别结果 ", index, ":", res.Result)
	}
	// ======================================================
	// 调用声学模型识别一段Wave音频序列
	result, err = sr.RecogniteSpeech(wave.GetRawSamples(), wave.FrameRate, wave.Channels, wave.SampleWidth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语音识别声学模型结果：", result.Result)
	// ======================================================
	// 调用语言模型1
	pinyinResult := []string{}
	for i := 0; i < len(result.Result.([]string)); i += 1 {
		pinyinResult = append(pinyinResult, result.Result.([]string)[i])
	}

	result, err = sr.RecogniteLanguage(pinyinResult)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语言模型结果：", result.Result)
	// ======================================================
	// 调用语言模型2
	sequencePinyin := []string{"ni3", "hao3", "a1"}
	result, err = sr.RecogniteLanguage(sequencePinyin)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语言模型结果：", result.Result)
	// ======================================================
	// 调用ASRT grpc接口流式识别
	wavChannel := make(chan *common.Wav, 5)
	recognitionResult := make(chan *common.AsrtAPIResponse, 5)
	sendFunction := func() {
		var index int
		for index = 0; index < 10; index += 1 {
			time.Sleep(2 * time.Second)
			wavChannel <- wave
		}
		close(wavChannel)
	}
	go sendFunction()
	var asrResult string
	var tmpAsrResult string
	recvFunction := func() {
		for value := range recognitionResult {
			fmt.Println("流式解码结果：", value.StatusCode, value.Result, value.StatucMesaage)
			if value.StatusCode == common.APIStatusCodeOK {
				tmpAsrResult = ""
				asrResult += value.Result.(string)
			} else if value.StatusCode == common.APIStatusCodePartOK {
				tmpAsrResult = value.Result.(string)
			}
			fmt.Println("语音识别文本：", asrResult+tmpAsrResult)
		}
	}
	go recvFunction()
	err = sr.(*sdk.GRPCSpeechRecognizer).RecogniteStream(wavChannel, recognitionResult)
	fmt.Println("流式识别完毕")
	if err != nil {
		fmt.Println(err)
	}
	close(recognitionResult)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
