import numpy as np
from PIL import Image
import os

def save_waveform(wave, name, folder="WaveForms/generated"):
    os.makedirs(folder, exist_ok=True)
    # Output: 1 non-white pixel per column, rest white (255)
    height = 128
    width = len(wave)
    img = np.full((height, width), 255, dtype=np.uint8)
    # Map waveform [-1,1] to y pixel (0=top, height-1=bottom)
    y = ((1 - wave) * (height - 1) / 2).astype(int)
    for x in range(width):
        img[y[x], x] = 0  # black pixel for waveform
    im = Image.fromarray(img, mode="L")
    im.save(f"{folder}/{name}.png")

# 33 Acoustic Bass: triangle with some harmonics
def acoustic_bass_wave(size=256):
    t = np.linspace(0, 1, size, endpoint=False)
    wave = 2 * np.abs(2 * t - 1) - 1  # triangle
    wave += 0.2 * np.sin(2 * np.pi * 3 * t)  # add 3rd harmonic
    return wave / np.max(np.abs(wave))

# 34 Electric Bass (finger): rounded square
def electric_bass_finger_wave(size=256):
    t = np.linspace(0, 1, size, endpoint=False)
    wave = np.sign(np.sin(2 * np.pi * t))
    wave = 0.7 * wave + 0.3 * np.sin(2 * np.pi * t)  # soften
    return wave / np.max(np.abs(wave))

# 35 Electric Bass (picked): sawtooth
def electric_bass_picked_wave(size=256):
    t = np.linspace(0, 1, size, endpoint=False)
    wave = 2 * (t - np.floor(t + 0.5))  # sawtooth
    return wave / np.max(np.abs(wave))

# 36 Fretless Bass: sine + triangle
def fretless_bass_wave(size=256):
    t = np.linspace(0, 1, size, endpoint=False)
    wave = 0.7 * np.sin(2 * np.pi * t) + 0.3 * (2 * np.abs(2 * t - 1) - 1)
    return wave / np.max(np.abs(wave))

# 37 Slap Bass 1: percussive, square + click
def slap_bass1_wave(size=256):
    t = np.linspace(0, 1, size, endpoint=False)
    wave = np.sign(np.sin(2 * np.pi * t))
    wave += 0.2 * np.random.uniform(-1, 1, size)  # add noise/click
    return wave / np.max(np.abs(wave))

# 38 Slap Bass 2: square + sawtooth
def slap_bass2_wave(size=256):
    t = np.linspace(0, 1, size, endpoint=False)
    wave = 0.5 * np.sign(np.sin(2 * np.pi * t)) + 0.5 * (2 * (t - np.floor(t + 0.5)))
    return wave / np.max(np.abs(wave))

# 39 Synth Bass 1: fat sawtooth
def synth_bass1_wave(size=256):
    t = np.linspace(0, 1, size, endpoint=False)
    wave = 2 * (t - np.floor(t + 0.5))
    wave += 0.5 * np.sin(2 * np.pi * 2 * t)  # add 2nd harmonic
    return wave / np.max(np.abs(wave))

# 40 Synth Bass 2: square + suboscillator
def synth_bass2_wave(size=256):
    t = np.linspace(0, 1, size, endpoint=False)
    wave = np.sign(np.sin(2 * np.pi * t))
    wave += 0.5 * np.sin(2 * np.pi * t / 2)  # suboscillator
    return wave / np.max(np.abs(wave))

if __name__ == "__main__":
    save_waveform(acoustic_bass_wave(), "acoustic_bass")
    save_waveform(electric_bass_finger_wave(), "electric_bass_finger")
    save_waveform(electric_bass_picked_wave(), "electric_bass_picked")
    save_waveform(fretless_bass_wave(), "fretless_bass")
    save_waveform(slap_bass1_wave(), "slap_bass_1")
    save_waveform(slap_bass2_wave(), "slap_bass_2")
    save_waveform(synth_bass1_wave(), "synth_bass_1")
    save_waveform(synth_bass2_wave(), "synth_bass_2")
