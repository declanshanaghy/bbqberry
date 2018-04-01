
export class HardwareConfig {
  vcc: number;
  analogMax: number;
  analogVoltsPerUnit: number;
  numLedPixels: number;
  probes: Probe[];
}

export class Probe {
  enabled: boolean;
  label: string;
  limits: TemperatureLimits;
}

export class TemperatureLimits {
  maxAbsCelsius: number;
  maxWarnCelsius: number;
  minAbsCelsius: number;
  minWarnCelsius: number;
  probeType: string;
}
