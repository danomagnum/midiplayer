import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "violin.png": lambda x: 0.7 * np.sin(2 * np.pi * x) + 0.3 * (2 * np.abs(2 * x - 1) - 1),  # Sine+triangle
    "viola.png": lambda x: 0.6 * np.sin(2 * np.pi * x) + 0.4 * (2 * np.abs(2 * x - 1) - 1),  # Sine+triangle, darker
    "cello.png": lambda x: 0.5 * np.sin(2 * np.pi * x) + 0.5 * (2 * np.abs(2 * x - 1) - 1),  # Sine+triangle, more triangle
    "contrabass.png": lambda x: 0.4 * np.sin(2 * np.pi * x) + 0.6 * (2 * np.abs(2 * x - 1) - 1),  # Sine+triangle, even darker
    "tremolo_strings.png": lambda x: 0.7 * np.sin(2 * np.pi * x) * np.abs(np.sin(16 * np.pi * x)),  # Sine with fast tremolo
    "pizzicato_strings.png": lambda x: np.sign(np.sin(2 * np.pi * x)) * np.exp(-4 * x),  # Plucked, percussive
    "orchestral_harp.png": lambda x: np.sin(2 * np.pi * x) * np.exp(-3 * x),  # Plucked, decaying
    "timpani.png": lambda x: np.sin(2 * np.pi * x) * (1 - 0.5 * x),  # Damped sine
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
