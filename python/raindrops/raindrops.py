sounds = {
    3: "Pling",
    5: "Plang",
    7: "Plong",
}


def convert(number):
    res = ""
    for key in sounds.keys():
        res += sounds[key] if number % key == 0 else ""

    return res if res else f"{number}"
