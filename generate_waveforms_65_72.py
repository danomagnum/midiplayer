import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "soprano_sax.png": lambda x: 0.7 * (2 * (x - np.floor(x + 0.5))) + 0.3 * np.sin(2 * np.pi * x),  # Sawtooth + sine
    "alto_sax.png": lambda x: 0.6 * (2 * (x - np.floor(x + 0.5))) + 0.4 * np.sin(2 * np.pi * x),  # Sawtooth + sine, darker
    "tenor_sax.png": lambda x: 0.5 * (2 * (x - np.floor(x + 0.5))) + 0.5 * np.sin(2 * np.pi * x),  # Sawtooth + sine, more sine
    "baritone_sax.png": lambda x: 0.4 * (2 * (x - np.floor(x + 0.5))) + 0.6 * np.sin(2 * np.pi * x),  # Sawtooth + sine, even darker
    "oboe.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * (2 * np.abs(2 * x - 1) - 1),  # Sine + triangle
    "english_horn.png": lambda x: 0.6 * np.sin(2 * np.pi * x) + 0.4 * (2 * np.abs(2 * x - 1) - 1),  # Sine + triangle, darker
    "bassoon.png": lambda x: 0.5 * np.sin(2 * np.pi * x) + 0.5 * (2 * np.abs(2 * x - 1) - 1),  # Sine + triangle, more triangle
    "clarinet.png": lambda x: 0.7 * np.sign(np.sin(2 * np.pi * x)) + 0.3 * np.sin(2 * np.pi * x),  # Square + sine
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
