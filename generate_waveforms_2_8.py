import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "bright_acoustic_piano.png": lambda x: np.sin(2 * np.pi * x) * 0.7 + np.sin(4 * np.pi * x) * 0.2,  # Brighter piano (sine+harmonics)
    "electric_grand_piano.png": lambda x: np.sign(np.sin(2 * np.pi * x)) * 0.7 + np.sin(2 * np.pi * x) * 0.3,  # Electric piano (square+sine)
    "honky_tonk_piano.png": lambda x: np.sin(2 * np.pi * x + 0.2 * np.sin(8 * np.pi * x)),  # Honky tonk (detuned sine)
    "electric_piano_1.png": lambda x: np.sin(2 * np.pi * x) * 0.7 + np.sin(8 * np.pi * x) * 0.2,  # Rhodes (sine+harmonics)
    "electric_piano_2.png": lambda x: np.sign(np.sin(2 * np.pi * x)) * 0.5 + np.sin(2 * np.pi * x) * 0.5,  # Wurlitzer (square+sine)
    "harpsichord.png": lambda x: 2 * (x - np.floor(x + 0.5)),  # Harpsichord (sawtooth)
    "clavinet.png": lambda x: np.sign(np.sin(2 * np.pi * x)) * 0.8,  # Clavinet (square)
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
