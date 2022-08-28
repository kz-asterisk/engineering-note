class Guitar:
  def __init__(self, tone):
    self.tone = tone
  
  def sounds(self):
    sound = f'{self.tone}♪♪♪'
    return sound
  
