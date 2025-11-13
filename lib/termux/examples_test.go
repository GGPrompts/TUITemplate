package termux_test

import (
	"fmt"
	"log"

	"github.com/yourname/TUITemplate/lib/termux"
)

// Example_basicNotification shows a simple notification.
func Example_basicNotification() {
	termux.Notify("Task Complete", "Your build finished successfully")
}

// Example_hapticFeedback shows how to add haptic feedback to UI interactions.
func Example_hapticFeedback() {
	// Quick vibration for button press
	termux.Vibrate(50)

	// Longer vibration for error
	termux.Vibrate(500)
}

// Example_toast shows quick status messages.
func Example_toast() {
	termux.Toast("File saved")
	termux.ToastLong("Processing complete - 15 files updated")
}

// Example_notificationWithButtons shows an interactive notification.
func Example_notificationWithButtons() {
	termux.Notify(
		"Pull Request Ready",
		"PR #123 has been created and is ready for review",
		termux.WithID("pr-123"),
		termux.WithPriority("high"),
		termux.WithButton("View PR", "termux-open-url https://github.com/user/repo/pull/123"),
		termux.WithButton("Copy URL", "bash -c 'echo https://github.com/user/repo/pull/123 | termux-clipboard-set'"),
		termux.WithVibrate("100,50,100"),
		termux.WithSound(),
	)
}

// Example_ongoingNotification shows a persistent notification for background tasks.
func Example_ongoingNotification() {
	// Show ongoing notification
	termux.Notify(
		"AI Worker Active",
		"Processing tasks in background...",
		termux.WithID("worker"),
		termux.WithOngoing(),
		termux.WithIcon("sync"),
	)

	// ... do work ...

	// Update notification
	termux.Notify(
		"AI Worker",
		"Completed 5 of 10 tasks",
		termux.WithID("worker"),
		termux.WithOngoing(),
	)

	// Remove when done
	termux.NotifyRemove("worker")
}

// Example_batteryAwareTask shows how to check battery before heavy operations.
func Example_batteryAwareTask() {
	battery, err := termux.GetBatteryStatus()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Battery: %d%% (%s)\n", battery.Percentage, battery.Status)

	// Skip heavy task if battery is low and not charging
	if battery.Percentage < 20 && battery.Status != "CHARGING" {
		termux.Toast("Low battery - skipping task")
		termux.Notify(
			"Task Skipped",
			fmt.Sprintf("Battery too low: %d%%", battery.Percentage),
			termux.WithPriority("low"),
		)
		return
	}

	// Acquire wake lock for long-running task
	termux.WakeLock()
	defer termux.WakeUnlock()

	// Do heavy work...
	fmt.Println("Processing...")
}

// Example_voiceInput shows voice-controlled workflows.
func Example_voiceInput() {
	termux.Toast("Listening...")

	text, err := termux.SpeechToText()
	if err != nil {
		log.Fatal(err)
	}

	termux.Speak(fmt.Sprintf("You said: %s", text))

	// Process voice command
	switch text {
	case "sync projects":
		termux.Vibrate(50)
		termux.Toast("Syncing projects...")
		// Run sync...
	case "check status":
		termux.Speak("All systems operational")
	}
}

// Example_confirmDialog shows user confirmation.
func Example_confirmDialog() {
	confirmed, err := termux.ConfirmDialog(
		"Approve PR?",
		"Merge pull request #123?",
	)
	if err != nil {
		log.Fatal(err)
	}

	if confirmed {
		termux.Vibrate(100)
		termux.Toast("PR approved")
		// Merge PR...
	} else {
		termux.Toast("Cancelled")
	}
}

// Example_textInput shows text input dialog.
func Example_textInput() {
	message, err := termux.TextDialog(
		"Commit Message",
		"Enter commit message:",
	)
	if err != nil {
		log.Fatal(err)
	}

	if message != "" {
		// Use commit message...
		termux.Toast("Committing changes...")
	}
}

// Example_radioSelection shows single-choice selection.
func Example_radioSelection() {
	action, err := termux.RadioDialog(
		"Choose Action",
		"Approve,Reject,Review,Cancel",
	)
	if err != nil {
		log.Fatal(err)
	}

	termux.Toast(fmt.Sprintf("Selected: %s", action))

	switch action {
	case "Approve":
		// Approve PR...
	case "Reject":
		// Reject PR...
	}
}

// Example_checkboxSelection shows multiple-choice selection.
func Example_checkboxSelection() {
	options, err := termux.CheckboxDialog(
		"Select Options",
		"Run Tests,Build,Deploy,Notify",
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, opt := range options {
		termux.Toast(fmt.Sprintf("Will run: %s", opt))
	}
}

// Example_location shows GPS location retrieval.
func Example_location() {
	termux.Toast("Getting location...")

	loc, err := termux.GetLocation()
	if err != nil {
		log.Fatal(err)
	}

	termux.Notify(
		"Location",
		fmt.Sprintf("%.4f, %.4f (±%.1fm)", loc.Latitude, loc.Longitude, loc.Accuracy),
		termux.WithButton("Copy", fmt.Sprintf("bash -c 'echo %.4f,%.4f | termux-clipboard-set'", loc.Latitude, loc.Longitude)),
	)
}

// Example_wifiAwareTask shows network-aware automation.
func Example_wifiAwareTask() {
	wifi, err := termux.GetWiFiConnectionInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to: %s (%d dBm)\n", wifi.SSID, wifi.RSSI)

	// Only run on trusted network
	if wifi.SSID != "HomeNetwork" {
		termux.Toast("Not on trusted network - skipping")
		return
	}

	// Safe to run automation
	fmt.Println("Running automation on home network...")
}

// Example_clipboard shows clipboard operations.
func Example_clipboard() {
	// Copy to clipboard
	url := "https://github.com/user/repo/pull/123"
	if err := termux.ClipboardSet(url); err != nil {
		log.Fatal(err)
	}
	termux.Toast("URL copied to clipboard")

	// Read from clipboard
	text, err := termux.ClipboardGet()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Clipboard:", text)
}

// Example_listInterface shows a TUI list with haptic feedback.
func Example_listInterface() {
	items := []string{"Item 1", "Item 2", "Item 3"}
	cursor := 0

	// Simulate navigation
	onUp := func() {
		if cursor > 0 {
			cursor--
			termux.Vibrate(30) // Quick haptic feedback
		}
	}

	onDown := func() {
		if cursor < len(items)-1 {
			cursor++
			termux.Vibrate(30)
		}
	}

	onSelect := func() {
		termux.Vibrate(50)
		termux.Toast(fmt.Sprintf("Selected: %s", items[cursor]))
		// Process selection...
	}

	// These would be called from your TUI's update function
	_ = onUp
	_ = onDown
	_ = onSelect
}

// Example_progressNotification shows updating notifications for progress tracking.
func Example_progressNotification() {
	tasks := []string{"Task 1", "Task 2", "Task 3", "Task 4", "Task 5"}

	for i, task := range tasks {
		// Update progress notification
		termux.Notify(
			"Processing Tasks",
			fmt.Sprintf("%d/%d: %s", i+1, len(tasks), task),
			termux.WithID("progress"),
			termux.WithOngoing(),
		)

		// Do work...
		fmt.Println("Processing", task)
	}

	// Final notification
	termux.Notify(
		"Complete",
		fmt.Sprintf("All %d tasks finished", len(tasks)),
		termux.WithID("progress"),
		termux.WithSound(),
		termux.WithVibrate("100,50,100,50,100"),
	)

	termux.Speak("All tasks complete")
}

// Example_sensorMonitoring shows reading device sensors.
func Example_sensorMonitoring() {
	// Check light level
	light, err := termux.GetSensor("light")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Light sensor: %v\n", light.Values)

	// Check proximity
	proximity, err := termux.GetSensor("proximity")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Proximity: %v\n", proximity.Values)

	// List all available sensors
	sensors, err := termux.ListSensors()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available sensors:", sensors)
}

// Example_backgroundWorker shows a complete background worker pattern.
func Example_backgroundWorker() {
	// Acquire wake lock
	termux.WakeLock()
	defer termux.WakeUnlock()

	// Check battery
	battery, _ := termux.GetBatteryStatus()
	if battery.Percentage < 20 && battery.Status != "CHARGING" {
		termux.Notify("Worker Skipped", "Battery too low", termux.WithPriority("low"))
		return
	}

	// Show ongoing notification
	termux.Notify(
		"Worker Active",
		"Processing tasks...",
		termux.WithID("worker"),
		termux.WithOngoing(),
	)

	// Do work
	successCount := 0
	failCount := 0

	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		termux.Toast(fmt.Sprintf("Processing: %s", task))

		// Simulate work...
		success := true // Replace with actual work

		if success {
			successCount++
			termux.Vibrate(50)
		} else {
			failCount++
			termux.Vibrate(200)
		}
	}

	// Final notification
	termux.Notify(
		"Worker Complete",
		fmt.Sprintf("%d succeeded, %d failed", successCount, failCount),
		termux.WithID("worker-done"),
		termux.WithPriority("high"),
		termux.WithSound(),
		termux.WithButton("View Results", "termux-open-url https://..."),
	)

	termux.Speak(fmt.Sprintf("All tasks complete. %d succeeded", successCount))
	termux.NotifyRemove("worker")
}

// Example_detection shows how to detect Termux environment.
func Example_detection() {
	if termux.IsTermux() {
		fmt.Println("Running on Termux")
		termux.Toast("Termux detected")
	} else {
		fmt.Println("Not running on Termux - using fallback behavior")
	}
}
