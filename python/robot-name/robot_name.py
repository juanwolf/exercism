import random


ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
NUMBERS = "0123456789"


class Robot(object):
    ROBOT_NAME_CHARS_NUMBER = 2
    ROBOT_NAME_NUM_NUMBERS = 3

    def __init__(self):
        self.reset()

    @staticmethod
    def _generate_name_section(collection, number_of_iteration):
        """_generate_name_section will return a random string of the size of number_of_iteration from the collection string"""
        res = ""
        for i in range(number_of_iteration):
            number_index = random.randint(0, len(collection) - 1)
            res += collection[number_index]
        return res

    def reset(self):
        random.seed()
        name = ""
        name += Robot._generate_name_section(ALPHABET, self.ROBOT_NAME_CHARS_NUMBER)
        name += Robot._generate_name_section(NUMBERS, self.ROBOT_NAME_NUM_NUMBERS)
        self.name = name
