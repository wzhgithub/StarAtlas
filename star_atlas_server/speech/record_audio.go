package speech

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
	wave "github.com/zenwerk/go-wave"
)

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
			httpClient()
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
