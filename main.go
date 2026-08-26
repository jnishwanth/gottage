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
	sceneFlag := flag.String("scene", "", "Scene to render with --demo (display name, or a unique substring of one)")
	weatherFlag := flag.String("weather", "", "Initial weather (clear, rain, snow, storm, fog)")
	timeFlag := flag.String("time", "", "Initial time of day (dawn, day, sunset, night, auto)")
	flag.Parse()

	// --scene only means something in demo mode. Accepting it silently in the
	// interactive path would leave the user staring at a scene they did not ask
	// for, with nothing to explain why.
	if *sceneFlag != "" && !*demoFlag {
		fmt.Fprintln(os.Stderr, "gottage: --scene requires --demo")
		os.Exit(2)
	}

	if *demoFlag {
		scene := art.NewCottageScene(74, 52)
		if *sceneFlag != "" {
			// Resolved through the registry, never through art.GetScene, which
			// answers an unregistered type with SceneRainyCity -- a demo that
			// quietly renders the wrong scene would lie in exactly the way this
			// flag exists to detect.
			st, err := art.SceneByName(*sceneFlag)
			if err != nil {
				fmt.Fprintf(os.Stderr, "gottage: %v\n\nAvailable scenes:\n", err)
				for _, name := range art.SceneNames() {
					fmt.Fprintf(os.Stderr, "  %s\n", name)
				}
				os.Exit(2)
			}
			scene.CurrentScene = st
		}
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
