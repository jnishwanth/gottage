package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gottage/pkg/art"
	"gottage/pkg/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	demoFlag := flag.Bool("demo", false, "Render a single frame of cozy bit-art to terminal and exit")
	weatherFlag := flag.String("weather", "", "Initial weather (clear, rain, snow, storm, fog)")
	timeFlag := flag.String("time", "", "Initial time of day (dawn, day, sunset, night, auto)")
	flag.Parse()

	if *demoFlag {
		scene := art.NewCottageScene(74, 68)
		if *weatherFlag != "" {
			switch *weatherFlag {
			case "clear":
				scene.Weather = art.WeatherClear
			case "rain":
				scene.Weather = art.WeatherRain
			case "snow":
				scene.Weather = art.WeatherSnow
			case "storm":
				scene.Weather = art.WeatherStorm
			case "fog":
				scene.Weather = art.WeatherFog
			}
		}
		scene.Update()
		fmt.Print(scene.Render())
		fmt.Println("Gottage — Virtual Cottage TUI (Demo Frame Rendered)")
		os.Exit(0)
	}

	model := ui.NewModel()

	// Apply CLI overrides if set
	if *weatherFlag != "" {
		switch *weatherFlag {
		case "clear":
			model.Scene.Weather = art.WeatherClear
		case "rain":
			model.Scene.Weather = art.WeatherRain
		case "snow":
			model.Scene.Weather = art.WeatherSnow
		case "storm":
			model.Scene.Weather = art.WeatherStorm
		case "fog":
			model.Scene.Weather = art.WeatherFog
		}
	}
	if *timeFlag != "" {
		switch *timeFlag {
		case "dawn":
			model.Scene.TimeOfDay = art.TimeDawn
		case "day":
			model.Scene.TimeOfDay = art.TimeDay
		case "sunset":
			model.Scene.TimeOfDay = art.TimeSunset
		case "night":
			model.Scene.TimeOfDay = art.TimeNight
		case "auto":
			model.Scene.TimeOfDay = art.TimeAuto
		}
	}

	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	// Ensure cleanup on exit
	defer func() {
		if model.Ambient != nil {
			model.Ambient.Cleanup()
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running Gottage: %v\n", err)
		os.Exit(1)
	}

	// Sleep briefly for cleanup
	time.Sleep(100 * time.Millisecond)
}
