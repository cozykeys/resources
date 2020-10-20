#!/usr/bin/env python3

# Takes raw image files from a phone as input and generates resized, renamed,
# and rotation-corrected image files as output ready for use on CozyKeys.xyz.
#
# Examples:
#
#     # Dry run
#     ./prepare_images.py -d -o ./temp -s 800x800,1600x1600 image1.jpg image2.jpg
#
#     # Live run
#     ./prepare_images.py -o ./temp -s 800x800,1600x1600 image1.jpg image2.jpg

from typing import List
import argparse
import errno
import fractions
import inspect
import os
import re
import subprocess

from PIL import Image  # type: ignore


class Dimensions:
    def __init__(self, width: int = 0, height: int = 0):
        self.width = width
        self.height = height

    def __str__(self):
        return "{width}x{height}".format(width=self.width, height=self.height)

    def get_aspect_ratio(self) -> fractions.Fraction:
        return fractions.Fraction(self.width, self.height)


def create_directory(path: str) -> None:
    try:
        os.makedirs(path)
    except OSError as e:
        if e.errno != errno.EEXIST:
            raise


def parse_target_dimensions(sizes_str: str) -> List[Dimensions]:
    sizes = []
    for size in sizes_str.split(","):
        if (m := re.match(r"^(\d+)x(\d+)$", size)) :
            width = int(m.group(1))
            height = int(m.group(2))
            sizes.append(Dimensions(width, height))
        else:
            raise Exception(
                "Invalid sizes argument, must be a comma-separated list in the form 800x800,1600x1600"
            )
    return sizes


def parse_args():
    p = argparse.ArgumentParser()
    p.add_argument(
        "-s",
        "--sizes",
        required=True,
        help="comma-separated list of sizes to generate (I.E. 800x800,1600x1600",
    )
    p.add_argument(
        "-o",
        "--output-dir",
        required=True,
        help="output directory to write resulting files to -- does not have to exist",
    )
    p.add_argument(
        "-d",
        "--dry-run",
        action="store_true",
        help="print the steps but don't execute them",
    )
    p.add_argument("images", nargs="*")
    return p.parse_args()


def main() -> int:
    args = parse_args()

    create_directory(args.output_dir)

    sizes = parse_target_dimensions(args.sizes)

    for img in args.images:
        print("Processing image {}:".format(img))
        if (m := re.match(r"(.*)\..*", img)) :
            name = m.group(1)
            im = Image.open(img)
            width, height = im.size
            aspect_ratio = fractions.Fraction(width, height)
            for size in sizes:
                if not aspect_ratio == size.get_aspect_ratio():
                    raise Exception(
                        "Aspect ratio of {} ({}) does not match target aspect ratio {}".format(
                            img, aspect_ratio, size.get_aspect_ratio()
                        )
                    )

                cmd = [
                    "convert",
                    img,
                    "-resize",
                    "{}".format(size),
                    "-auto-orient",
                    "{}/{}_{}.png".format(args.output_dir, name, size),
                ]

                print("    {}".format(" ".join(cmd)))

                if not args.dry_run:
                    subprocess.run(cmd)
        else:
            raise Exception("Invalid image path {}".format(img))

    return 0


if __name__ == "__main__":
    exit(main())
