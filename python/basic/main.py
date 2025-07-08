def toString(nbr):
    return str(nbr)


def isString(string):
    if isinstance(string, str):
        return "La variable est une chaîne de caractères"


if __name__ == "__main__":
    number = 15
    print("Le nombre est " + toString(number))
    print(isString(number))
    phrase = "Bonjour tout le monde."
    nouvelle_phrase = "Bonsoir" + phrase[7:]
    print(nouvelle_phrase)
