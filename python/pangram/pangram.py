ALPHABET = "abcdefghijklmnopqrstuvwxyz"


def is_pangram(sentence):
    letter_encountered = {letter: False for letter in ALPHABET}
    sentence = sentence.lower()
    for letter in sentence:
        if letter in letter_encountered:
            letter_encountered[letter] = True

    for encountered in letter_encountered.values():
        if not encountered:
            return encountered

    return True
