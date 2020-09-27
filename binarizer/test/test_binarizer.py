from unittest import TestCase
import enum
import io
from typing import List, Callable

import matplotlib
import matplotlib.pyplot as plt
import numpy
from skimage import img_as_ubyte
from skimage.color import rgb2gray

from skimage.data import page, chelsea
from skimage.filters import (threshold_otsu, threshold_niblack,
                             threshold_sauvola)

from app.binarizer import Binarizer


class TestBinarizer(TestCase):
    def test_binarize(self):
        b = Binarizer()
        binary_sauvola = b.binarize([page()], Method.SAUVOLA)
        plt.imshow(binary_sauvola[0], cmap=plt.cm.gray)
        plt.show()
