import type { PrinterStatus } from "@jamesgopsill/octoprint-client";

import TemperatureListItem from "./TemperatureListItem";

type TemperatureListProps = {
  temperatures: Omit<PrinterStatus["temperature"], "history">;
};

function TemperatureList({ temperatures }: TemperatureListProps) {
  return Object.keys(temperatures).map((k) => <TemperatureListItem key={k} name={k.replace("tool", "tool ")} temperature={temperatures[k as keyof typeof temperatures]} />)
}

export default TemperatureList;