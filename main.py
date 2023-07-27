from PIL import Image
import easygui


paths = easygui.fileopenbox(multiple=True, title="Select Images")
directory = easygui.diropenbox()

count = 0
for path in paths:
    count += 1
    image = Image.open(path)

    if image.size[0] > image.size[1]:
        new = Image.new(mode="RGBA", size=(image.size[0],  image.size[0]), color="white")
        new.paste(image, (0, (int((image.size[0] - image.size[1])/2))))
    else:
        new = Image.new(mode="RGBA", size=(image.size[1], image.size[1]), color="white")
        new.paste(image, ((int((image.size[1] - image.size[0])/2)), 0))

    new.save(f"{directory}\IMG_{count}.png")

