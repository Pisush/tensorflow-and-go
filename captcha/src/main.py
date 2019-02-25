import random
from captcha.image import ImageCaptcha
def random_string(length):
    return ''.join(random.choice("ABCDEFGHJKLMNPQRSTUVWXYZ23456789") for _ in range(length))
captcha_gen = ImageCaptcha(width=160, height=60, font_sizes=[42])
for i in xrange(5):
  captcha_string = random_string(6)
  captcha_gen.write(captcha_string, captcha_string+'.png')
