import numpy as np
from PIL import Image, ImageDraw
import os

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

def organ_wave(x, harmonics):
    y = np.zeros_like(x)
    for i, amp in enumerate(harmonics):
        y += amp * np.sin(2 * np.pi * (i+1) * x)
    return y / np.max(np.abs(y))

waveforms = {
    "drawbar_organ.png": lambda x: organ_wave(x, [1, 0.7, 0.5, 0.3, 0.2, 0.1]),  # Drawbar organ
    "percussive_organ.png": lambda x: organ_wave(x, [1, 0.5, 0.2]),  # Percussive organ
    "rock_organ.png": lambda x: organ_wave(x, [1, 0.8, 0.6, 0.4, 0.2]),  # Rock organ
    "church_organ.png": lambda x: organ_wave(x, [1, 0.6, 0.4, 0.3, 0.2, 0.1]),  # Church organ
    "reed_organ.png": lambda x: organ_wave(x, [1, 0.5, 0.3]),  # Reed organ
    "accordion.png": lambda x: organ_wave(x, [1, 0.7, 0.4, 0.2]),  # Accordion
    "harmonica.png": lambda x: organ_wave(x, [1, 0.8, 0.5, 0.2]),  # Harmonica
    "bandoneon.png": lambda x: organ_wave(x, [1, 0.6, 0.3, 0.2]),  # Bandoneon
}

def draw_waveform(wave_fn, filename):
    x = np.linspace(0, 1, WIDTH)
    y = wave_fn(x)
    y = np.clip(y, -1, 1)
    img = Image.new("L", (WIDTH, HEIGHT), 255)
    draw = ImageDraw.Draw(img)
    points = [(i, int(MID_Y - y[i] * (HEIGHT // 2 - 10))) for i in range(WIDTH)]
    draw.line(points, fill=0, width=2)
    os.makedirs("WaveForms/generated", exist_ok=True)
    img.save(f"WaveForms/generated/{filename}")

if __name__ == "__main__":
    for fname, fn in waveforms.items():
        draw_waveform(fn, fname)
