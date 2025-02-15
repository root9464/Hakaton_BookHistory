export const project = (lat: number, lon: number, zoom: number) => {
  const scale = 1 << zoom;
  const worldSize = 256 * scale;
  const x = ((lon + 180) * worldSize) / 360;
  const siny = Math.sin((lat * Math.PI) / 180);
  const y = (0.5 - Math.log((1 + siny) / (1 - siny)) / (4 * Math.PI)) * worldSize;
  return { x, y };
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
