import type { TemperatureData } from "../types"

function TemperatureListItem({ temperature, name }: { temperature: TemperatureData, name: string }) {
  return (
    <div id={name} className="mb-3">
      <h5 className="text-capitalize">{name.replace("-", " ")}
      </h5>
      <div>Actual: {temperature.Actual} &deg;C
      </div>
      <div>Target: {temperature.Target} &deg;C
      </div>
    </div>
  )
}

export default TemperatureListItem