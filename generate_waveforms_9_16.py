import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

def bell_wave(x, overtone_weights):
    y = np.zeros_like(x)
    for i, w in enumerate(overtone_weights):
        y += w * np.sin(2 * np.pi * (i+1) * x)
    return y / np.max(np.abs(y))

waveforms = {
    "celesta.png": lambda x: bell_wave(x, [1, 0.5, 0.2, 0.1]),  # Bell-like, soft
    "glockenspiel.png": lambda x: bell_wave(x, [1, 0.7, 0.3, 0.15, 0.1]),  # Bright bell
    "music_box.png": lambda x: bell_wave(x, [1, 0.6, 0.3, 0.2, 0.1]),  # Plucked bell
    "vibraphone.png": lambda x: bell_wave(x, [1, 0.5, 0.3, 0.2, 0.1]),  # Mellow bell
    "marimba.png": lambda x: np.sin(2 * np.pi * x) * 0.7 + np.sin(6 * np.pi * x) * 0.2,  # Wood, mellow
    "xylophone.png": lambda x: np.sign(np.sin(2 * np.pi * x)) * 0.7 + np.sin(6 * np.pi * x) * 0.2,  # Bright, woody
    "tubular_bells.png": lambda x: bell_wave(x, [1, 0.8, 0.4, 0.2, 0.1]),  # Deep bell
    "dulcimer.png": lambda x: np.sin(2 * np.pi * x) * 0.7 + np.sin(8 * np.pi * x) * 0.2,  # Plucked string
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
