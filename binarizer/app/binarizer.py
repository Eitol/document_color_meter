import enum
import os
import string
import random
import string

from pathlib import Path
from typing import List, Callable
from io import BytesIO
import numpy
from PIL import Image as pim
from imageio.core import Image
from skimage.color import rgb2gray
import matplotlib.pyplot as plt
from skimage.filters import threshold_sauvola, threshold_niblack, threshold_otsu
from skimage.io._plugins.pil_plugin import pil_to_ndarray


class Method(enum.Enum):
    UNSET = 0
    SAUVOLA = 1
    NIBLACK = 2
    OTSU = 3


DEFAULT_OUT_PATH = os.path.join(str(Path.home()), "out")


class Binarizer(object):
    out_path = ""
    
    @classmethod
    def binarize(cls, images: List[bytes], method: Method) -> (List[bytes], str):
        if method == Method.SAUVOLA or method == Method.UNSET:
            return cls._binarize(images, lambda x: threshold_sauvola(x, window_size=25))
        elif method == Method.NIBLACK:
            
            return cls._binarize(images, lambda x: threshold_niblack(x, window_size=25, k=0.8))
        else:
            return cls._binarize(images, lambda x: threshold_otsu(x))
    
    @classmethod
    def _get_random_name(cls):
        letters = string.ascii_lowercase
        r = ''.join(random.choice(letters) for i in range(10))
        d = os.path.join(cls._get_out_path(), r)
        os.makedirs(d, exist_ok=True)
        return d
    
    @classmethod
    def _binarize(cls, images: List[bytes], method: Callable[[numpy.ndarray], numpy.ndarray]) -> (List[bytes], str):
        out_path = cls._get_random_name()
        out: List[bytes] = []
        count = 1
        for image_bytes in images:
            buf = BytesIO(image_bytes)
            im = pim.open(buf)
            image = Image(pil_to_ndarray(im))
            grayscale = rgb2gray(image)
            # thresh: numpy.ndarray = method(grayscale)
            # binary: numpy.ndarray = grayscale > thresh
            img_bytes: bytes = cls._binary_image_to_bytes(grayscale)
            with open(os.path.join(out_path, f"_{count}.jpg"), "wb") as f:
                f.write(img_bytes)
            out.append(img_bytes)
            buf.close()
            count += 1
        return out, out_path
    
    @classmethod
    def _get_out_path(cls) -> str:
        if cls.out_path != "":
            return cls.out_path
        op = os.getenv("OUT_PATH")
        if op is None or op == "":
            op = DEFAULT_OUT_PATH
        if not os.path.exists(op):
            os.makedirs(op, exist_ok=True)
        cls.out_path = op
        return op
    
    @classmethod
    def _binary_image_to_bytes(cls, binary: numpy.ndarray) -> bytes:
        shape = numpy.shape(binary)[0:2][::-1]
        size = [float(i) / 100 for i in shape]
        fig = plt.figure()
        fig.set_size_inches(size)
        ax = plt.Axes(fig, [0, 0, 1, 1])
        ax.set_axis_off()
        fig.add_axes(ax)
        ax.imshow(binary, cmap="gray")
        buf = BytesIO()
        plt.savefig(buf, format='jpeg')
        plt.close()
        buf.seek(0)
        return buf.read()
