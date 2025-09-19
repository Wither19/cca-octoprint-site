import type { PrinterStatus } from "@jamesgopsill/octoprint-client";

export function printerStateColors(s: PrinterStatus["state"]["flags"]) {
	let color = "secondary";

	if (s) {
		if (s["ready"] && s["sdReady"]) {
			color = "success";
		} else if (s["paused"] || s["printing"] || s["cancelling"]) {
			color = "warning";
		} else if (s["error"] || s["closedOnError"]) {
			color = "danger";
		}
	}

	return color;
}