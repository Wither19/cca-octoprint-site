export type TemperatureData = {
  "actual": number;
  "target": number;
  "offset": number;
};

export type TemperatureObject = {
  "bed": TemperatureData;
  "tool0": TemperatureData;
  "tool1": TemperatureData;
  "tool2": TemperatureData;
  "tool3": TemperatureData;
  "tool4": TemperatureData;
  "history"?: TemperatureHistory[];
};

export type TemperatureHistory = {
  "time": number;
  "tool": { [key: string]: TemperatureData };
};

export type PrinterStateFlags = {
  "operational": boolean;
  "paused": boolean;
  "printing": boolean;
  "cancelling": boolean;
  "pausing": boolean;
  "sdready": boolean;
  "error": boolean;
  "ready": boolean;
  "closedorerror": boolean;
};

export type PrinterState = {
  "flags": PrinterStateFlags;
  "text": string;
  "error": string;
  "temperature": TemperatureObject;
};
