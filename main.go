package main

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"slices"
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

	// Whether --scene was passed, not whether its value is non-empty. `--scene
	// "$SCENE"` with SCENE unset arrives here as the empty string, and reading
	// that as "no --scene given" renders whatever the default happens to be --
	// the silent-wrong-scene failure this flag exists to detect. An empty value
	// is a request that cannot be satisfied; SceneByName says so.
	sceneRequested := flagPassed(flag.CommandLine, "scene")

	// --scene only means something in demo mode. Accepting it silently in the
	// interactive path would leave the user staring at a scene they did not ask
	// for, with nothing to explain why.
	if sceneRequested && !*demoFlag {
		fmt.Fprintln(os.Stderr, "gottage: --scene requires --demo")
		os.Exit(2)
	}

	// Resolved once, for both paths. These used to be two switch statements with
	// no default arm, so a value no arm claimed -- and an empty one, for the same
	// reason --scene now keys off flagPassed -- was accepted and dropped, and
	// --time never reached the demo path at all.
	weather := art.WeatherClear
	weatherRequested := flagPassed(flag.CommandLine, "weather")
	if weatherRequested {
		w, err := parseWeather(*weatherFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gottage: %v\n", err)
			os.Exit(2)
		}
		weather = w
	}

	timeOfDay := art.TimeAuto
	timeRequested := flagPassed(flag.CommandLine, "time")
	if timeRequested {
		tod, err := parseTimeOfDay(*timeFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gottage: %v\n", err)
			os.Exit(2)
		}
		timeOfDay = tod
	}

	if *demoFlag {
		scene := art.NewCottageScene(74, 52)
		if sceneRequested {
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
		if weatherRequested {
			scene.Weather = weather
		}
		if timeRequested {
			scene.TimeOfDay = timeOfDay
		}
		scene.Update()
		fmt.Print(scene.Render())

		// A frame drawn from the navy fallback is not a rendered scene, and this
		// is the command the phase doc uses to prove the binary resolves assets
		// off the repo root. Reporting it costs nothing here: --demo is one-shot
		// and never enters bubbletea's alt screen, which is what the in-memory
		// -only rule in pkg/art protects.
		if errs := art.AssetErrors(); len(errs) != 0 {
			fmt.Fprintf(os.Stderr, "gottage: %d asset(s) failed to load; the frame above is the fallback, not the scene:\n", len(errs))
			for _, name := range slices.Sorted(maps.Keys(errs)) {
				fmt.Fprintf(os.Stderr, "  %s: %v\n", name, errs[name])
			}
			os.Exit(1)
		}

		fmt.Println("Gottage — Virtual Cottage TUI (Demo Frame Rendered)")
		os.Exit(0)
	}

	model := ui.NewModel()

	// Apply CLI overrides if set
	if weatherRequested {
		model.Scene.Weather = weather
	}
	if timeRequested {
		model.Scene.TimeOfDay = timeOfDay
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

// parseWeather and parseTimeOfDay reject anything they do not recognise rather
// than falling off the end of a switch, so a typo is an exit-2 diagnostic instead
// of a flag that was accepted and had no effect.
func parseWeather(s string) (art.WeatherMode, error) {
	switch s {
	case "clear":
		return art.WeatherClear, nil
	case "rain":
		return art.WeatherRain, nil
	case "snow":
		return art.WeatherSnow, nil
	case "storm":
		return art.WeatherStorm, nil
	case "fog":
		return art.WeatherFog, nil
	}
	return art.WeatherClear, fmt.Errorf("unknown --weather %q (want clear, rain, snow, storm or fog)", s)
}

func parseTimeOfDay(s string) (art.TimeOfDay, error) {
	switch s {
	case "dawn":
		return art.TimeDawn, nil
	case "day":
		return art.TimeDay, nil
	case "sunset":
		return art.TimeSunset, nil
	case "night":
		return art.TimeNight, nil
	case "auto":
		return art.TimeAuto, nil
	}
	return art.TimeAuto, fmt.Errorf("unknown --time %q (want dawn, day, sunset, night or auto)", s)
}

// flagPassed reports whether name appeared on the command line. FlagSet.Visit
// walks only the flags actually set, which is the distinction a flag whose zero
// value is also a legal-looking value cannot draw from its value alone.
func flagPassed(fs *flag.FlagSet, name string) bool {
	passed := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == name {
			passed = true
		}
	})
	return passed
}
