import numpy as np
from PIL import Image, ImageDraw
import os

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

def plucked_string_wave(x, harmonics, decay=1.0):
    y = np.zeros_like(x)
    for i, amp in enumerate(harmonics):
        y += amp * np.sin(2 * np.pi * (i+1) * x) * (decay ** i)
    return y / np.max(np.abs(y))

def muted_wave(x):
    # Muted: short, percussive, high harmonics
    return plucked_string_wave(x, [1, 0.7, 0.5, 0.3, 0.2], decay=0.5)

def overdriven_wave(x):
    # Overdriven: clipped sine
    return np.clip(2 * np.sin(2 * np.pi * x), -0.7, 0.7)

def distortion_wave(x):
    # Distortion: hard clipped
    return np.sign(np.sin(2 * np.pi * x)) * 0.7

def harmonics_wave(x):
    # Guitar harmonics: high, bell-like
    return plucked_string_wave(x, [0, 1, 0, 0.7, 0, 0.5, 0, 0.3], decay=0.7)

waveforms = {
    "acoustic_guitar_nylon.png": lambda x: plucked_string_wave(x, [1, 0.6, 0.3, 0.15]),
    "acoustic_guitar_steel.png": lambda x: plucked_string_wave(x, [1, 0.7, 0.4, 0.2]),
    "electric_guitar_jazz.png": lambda x: plucked_string_wave(x, [1, 0.5, 0.2, 0.1]),
    "electric_guitar_clean.png": lambda x: plucked_string_wave(x, [1, 0.6, 0.3, 0.1]),
    "electric_guitar_muted.png": muted_wave,
    "overdriven_guitar.png": overdriven_wave,
    "distortion_guitar.png": distortion_wave,
    "guitar_harmonics.png": harmonics_wave,
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
