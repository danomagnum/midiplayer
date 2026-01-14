import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "new_age.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.sin(4 * np.pi * x),
    "warm.png": lambda x: 0.8 * np.sin(2 * np.pi * x) + 0.2 * np.sin(3 * np.pi * x),
    "polysynth.png": lambda x: 0.6 * np.sin(2 * np.pi * x) + 0.4 * (2 * (x - np.floor(x + 0.5))),
    "choir.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.sin(6 * np.pi * x),
    "bowed.png": lambda x: 0.5 * np.sin(2 * np.pi * x) + 0.5 * (2 * np.abs(2 * x - 1) - 1),
    "metallic.png": lambda x: 0.5 * np.sin(2 * np.pi * x) + 0.5 * np.sin(8 * np.pi * x),
    "halo.png": lambda x: 0.6 * np.sin(2 * np.pi * x) + 0.4 * np.sin(5 * np.pi * x),
    "sweep.png": lambda x: 0.7 * np.sin(2 * np.pi * x) * np.exp(-2 * x),
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
