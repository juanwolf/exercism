ALPHABET = "abcdefghijklmnopqrstuvwxyz"


def is_pangram(sentence):
    letter_found = {letter: False for letter in ALPHABET}
    sentence = sentence.lower()
    for letter in sentence:
        if letter in letter_found:
            letter_found[letter] = True

    for found in letter_found.values():
        if not found:
            return found

    return True
