export const project = (lat: number, lon: number, zoom: number) => {
  const scale = 1 << zoom;
  const worldSize = 256 * scale;
  const x = ((lon + 180) * worldSize) / 360;
  const siny = Math.sin((lat * Math.PI) / 180);
  const y = (0.5 - Math.log((1 + siny) / (1 - siny)) / (4 * Math.PI)) * worldSize;
  return { x, y };
};

export const unproject = (x: number, y: number, zoom: number) => {
  const scale = 1 << zoom;
  const worldSize = 256 * scale;

  const normalizedX = ((x % worldSize) + worldSize) % worldSize;
  const lon = (normalizedX * 360) / worldSize - 180;

  const yRelative = 0.5 - y / worldSize;
  const latRad = Math.atan(Math.sinh(yRelative * 2 * Math.PI));
  const lat = (latRad * 180) / Math.PI;

  return { lat, lon };
};

export const toEPSG3857Direct = (worldX: number, worldY: number, zoom: number) => {
  const scale = 1 << zoom;
  const worldSize = 256 * scale;
  const earthCircumference = 40075016.68557849;

  return {
    x: (worldX / worldSize) * earthCircumference - earthCircumference / 2,
    y: earthCircumference / 2 - (worldY / worldSize) * earthCircumference,
  };
};

//zombie code
export const toEPSG3857 = (lat: number, lon: number) => {
  const x = (lon * 20037508.34) / 180;
  const y = (Math.log(Math.tan(((90 + lat) * Math.PI) / 360)) * 20037508.34) / Math.PI;
  return { x, y };
};
