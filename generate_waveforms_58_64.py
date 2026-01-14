import numpy as np
from PIL import Image, ImageDraw

WIDTH, HEIGHT = 256, 200
MID_Y = HEIGHT // 2

waveforms = {
    "trombone.png": lambda x: 0.7 * (2 * (x - np.floor(x + 0.5))) + 0.3 * np.sin(2 * np.pi * x),  # Sawtooth + sine
    "tuba.png": lambda x: 0.5 * (2 * (x - np.floor(x + 0.5))) + 0.5 * np.sin(2 * np.pi * x),  # Sawtooth + sine, darker
    # TODO: this muted trumpet sounds terrible.
    "muted_trumpet.png": lambda x: 0.6 * np.sin(2 * np.pi * x) * np.abs(np.sin(8 * np.pi * x)),  # Sine with fast tremolo (mute)
    "french_horn.png": lambda x: 0.5 * np.sin(2 * np.pi * x) + 0.5 * (2 * np.abs(2 * x - 1) - 1),  # Sine + triangle
    "brass_section.png": lambda x: 0.8 * (2 * (x - np.floor(x + 0.5))) + 0.2 * np.sin(2 * np.pi * x),  # Brighter sawtooth
    "synth_brass_1.png": lambda x: 1.2 * (2 * (x - np.floor(x + 0.5))),  # Fat sawtooth
    "synth_brass_2.png": lambda x: 0.7 * np.sign(np.sin(2 * np.pi * x)) + 0.3 * (2 * (x - np.floor(x + 0.5))),  # Square/saw hybrid
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
