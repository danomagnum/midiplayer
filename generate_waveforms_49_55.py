import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "string_ensemble_1.png": lambda x: 2 * (x - np.floor(x + 0.5)),  # Sawtooth
    "string_ensemble_2.png": lambda x: 1.5 * (2 * (x - np.floor(x + 0.5))),  # Brighter sawtooth
    "synthstrings_1.png": lambda x: 1.2 * (2 * (x - np.floor(x + 0.5))),  # Synthetic saw
    "synthstrings_2.png": lambda x: 0.7 * np.sign(np.sin(2 * np.pi * x)) + 0.3 * (2 * (x - np.floor(x + 0.5))),  # Square/saw hybrid
    "choir_aahs.png": lambda x: np.sin(2 * np.pi * x) * 0.7 + np.sin(4 * np.pi * x) * 0.2,  # Smooth sine/triangle
    "voice_oohs.png": lambda x: np.sin(2 * np.pi * x) * 0.5 + np.sin(4 * np.pi * x) * 0.1,  # Softer sine/triangle
    "synth_voice.png": lambda x: np.abs(2 * (x - np.floor(x + 0.5))) * 0.7,  # Triangle
    "orchestra_hit.png": lambda x: np.sign(np.sin(2 * np.pi * x * 2)) * np.exp(-4 * x),  # Percussive, sharp
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
