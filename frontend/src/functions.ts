import type { PrinterStateFlags } from "./types"

export function printerStateColors(s: PrinterStateFlags) {
	let color: string

	if (s.Ready && s.SDReady) {
		color = "success"
	} else if (s.Paused || s.Printing || s.Cancelling || s.Pausing) {
		color = "warning"
	} else if (s.Error || s.ClosedOrError) {
		color = "danger"
	} else {
		color = "secondary"
	}

	return color
}