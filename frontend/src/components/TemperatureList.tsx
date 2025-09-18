import type { TemperatureObject } from "../types";

import TemperatureListItem from "./TemperatureListItem";

type TemperatureListProps = {
  temperatures: Omit<TemperatureObject, "history">;
};

function TemperatureList({ temperatures }: TemperatureListProps) {
  const existingKeys = Object.keys(temperatures);
  return existingKeys.map((k) => <TemperatureListItem key={k} name={k} temperature={temperatures[k as keyof typeof temperatures]} />)
}

export default TemperatureList;