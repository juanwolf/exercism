def two_fer(name="you"):
    # Setting name to "you" if name is undefined or empty
    if not name:
        name = "you"

    return "One for %s, one for me." % name
