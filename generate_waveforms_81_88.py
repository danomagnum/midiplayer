import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "square_wave.png": lambda x: np.sign(np.sin(2 * np.pi * x)),
    "sawtooth_wave.png": lambda x: 2 * (x - np.floor(x + 0.5)),
    "calliope.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.sin(6 * np.pi * x),
    "chiff.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.random.uniform(-0.5, 0.5, x.shape),
    "charang.png": lambda x: 0.6 * (2 * (x - np.floor(x + 0.5))) + 0.4 * np.sin(2 * np.pi * x),
    "voice.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * np.sin(4 * np.pi * x),
    "fifths.png": lambda x: np.sin(2 * np.pi * x) + 0.5 * np.sin(3 * np.pi * x),
    "bass_lead.png": lambda x: 0.7 * (2 * (x - np.floor(x + 0.5))) + 0.3 * np.sign(np.sin(2 * np.pi * x)),
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
