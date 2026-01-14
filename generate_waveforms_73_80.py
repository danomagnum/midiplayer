import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "piccolo.png": lambda x: 0.8 * np.sin(2 * np.pi * x) + 0.2 * np.sin(4 * np.pi * x),  # Bright sine
    "flute.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.sin(3 * np.pi * x),  # Sine + 3rd harmonic
    "recorder.png": lambda x: 0.6 * np.sin(2 * np.pi * x) + 0.4 * np.sin(5 * np.pi * x),  # Sine + 5th harmonic
    "pan_flute.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.exp(-3 * x),  # Sine + decay
    "blown_bottle.png": lambda x: 0.6 * np.sin(2 * np.pi * x) + 0.4 * np.random.uniform(-0.2, 0.2, x.shape),  # Sine + noise
    "shakuhachi.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.sin(6 * np.pi * x),  # Sine + 6th harmonic
    "whistle.png": lambda x: 0.9 * np.sin(2 * np.pi * x) + 0.1 * np.sin(8 * np.pi * x),  # Pure sine, slight overtone
    "ocarina.png": lambda x: 0.8 * np.sin(2 * np.pi * x) + 0.2 * np.sin(7 * np.pi * x),  # Sine + 7th harmonic
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
