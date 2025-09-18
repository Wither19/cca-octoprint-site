import type { PrinterStateFlags } from "./types"

export function printerStateColors(s: PrinterStateFlags) {
	let color = "secondary";

	if (s) {
		if (s["ready"] && s["sdready"]) {
			color = "success";
		} else if (s["paused"] || s["printing"] || s["cancelling"] || s["pausing"]) {
			color = "warning";
		} else if (s["error"] || s["closedorerror"]) {
			color = "danger";
		}
	}

	return color;
}