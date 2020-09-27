import os
from typing import Dict
from unittest import TestCase

from app.pdf_to_image_converter import PdfToImageConverter

test__data_path = 'testdata'

test_files: Dict[str, int] = {
    os.path.join(test__data_path, 'test.pdf'): 2,
    os.path.join(test__data_path, 'test_pdf_a.pdf'): 1
}


class TestPdfToImageConverter(TestCase):
    def test_convert(self):
        for path, page_count in test_files.items():
            with open(path, 'rb') as f:
                pdf_file = f.read()
                result = PdfToImageConverter.convert(pdf_file)
                self.assertEqual(len(result), page_count)
