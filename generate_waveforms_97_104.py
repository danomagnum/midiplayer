import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "fx_rain.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.random.uniform(-1, 1, x.shape),
    "fx_soundtrack.png": lambda x: 0.6 * np.sin(2 * np.pi * x) + 0.4 * np.sin(8 * np.pi * x),
    "fx_crystal.png": lambda x: 0.5 * np.sin(2 * np.pi * x) + 0.5 * np.abs(np.sin(12 * np.pi * x)),
    "fx_atmosphere.png": lambda x: 0.7 * np.sin(2 * np.pi * x) * np.exp(-2 * x),
    "fx_brightness.png": lambda x: 0.8 * np.sin(2 * np.pi * x) + 0.2 * np.sin(16 * np.pi * x),
    "fx_goblins.png": lambda x: 0.6 * np.sin(2 * np.pi * x) + 0.4 * np.sign(np.sin(6 * np.pi * x)),
    "fx_echo_drops.png": lambda x: 0.7 * np.sin(2 * np.pi * x) * np.abs(np.sin(10 * np.pi * x)),
    "fx_star_theme.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.sin(5 * np.pi * x),
}

def draw_waveform(wave_fn, filename):
    x = np.linspace(0, 1, WIDTH)
    y = wave_fn(x)
    y = np.clip(y, -1, 1)
    img = Image.new("L", (WIDTH, HEIGHT), 255)
    draw = ImageDraw.Draw(img)
    points = [(i, int(MID_Y - y[i] * (HEIGHT // 2 - 10))) for i in range(WIDTH)]
    draw.line(points, fill=0, width=2)
    img.save(f"WaveForms/generated/{filename}")

if __name__ == "__main__":
    for fname, fn in waveforms.items():
        draw_waveform(fn, fname)
