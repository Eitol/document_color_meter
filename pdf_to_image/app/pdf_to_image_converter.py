import abc
import enum
import io
import gc
from typing import List

import PIL
from pdf2image import convert_from_bytes


class ImageFormat(enum.Enum):
    UNSET = "unset"
    UNKNOWN = "unknown"
    JPEG = "jpeg"
    PNG = "png"


class FileToImageConverter(metaclass=abc.ABCMeta):
    
    @classmethod
    @abc.abstractmethod
    def convert(cls, file: bytes, fmt: ImageFormat) -> List[bytes]:
        pass


_WRITE_IN_FILE = False


class PdfToImageConverter(FileToImageConverter):
    
    @staticmethod
    def _image_to_byte_array(image: PIL.Image):
        with io.BytesIO() as img_byte_arr:
            image.save(img_byte_arr, format=image.format)
            return img_byte_arr.getvalue()
    
    @classmethod
    def convert(cls, file: bytes, fmt: ImageFormat = ImageFormat.JPEG) -> List[bytes]:
        """
        It is responsible for converting a PDF document to images
        :param file: Binary of the image
        :param fmt: Format of the binary
        :return: An array of images bytes.
        """
        images = convert_from_bytes(file, fmt=fmt.value, dpi=100)
        out = []
        for image in images:
            out.append(cls._image_to_byte_array(image))
        return out
