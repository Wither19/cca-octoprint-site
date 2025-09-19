import type { PrinterStatus } from "@jamesgopsill/octoprint-client"

function TemperatureListItem({ temperature, name }: { temperature: PrinterStatus["temperature"][keyof PrinterStatus["temperature"]], name: string }) {
  return (
    <div id={name} className="mb-3">
      <h3 className="text-capitalize">{name.replace("-", " ")}</h3>
      <div>Actual: {temperature!["actual"] ?? 0} &deg;C</div>
      <div>Target: {temperature!["target"] ?? 0} &deg;C</div>
    </div>
  )
}

export default TemperatureListItem