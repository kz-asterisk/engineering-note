class Monster:
    def __init__(self, name, hitpoint, atack):
        self.name = name
        self.hitpoint = hitpoint
        self.atack = atack
        self.nickname = self.name + "_" + "chan"
        if len(self.name) / 2 == 0:
            self.roar = "Ugaaah"
        else:
            self.roar = "Shaaah"

    def decribe(self):
        desc = f'name = {self.name} hitpoint = {self.hitpoint} atack = {self.atack} roar = {self.roar}'
        return desc
