from PIL import Image

#Read the two images
image1 = Image.open('generatedOR/baby_0.png')
image1.show()

image2 = Image.open('generatedOR/baby_1.png')
image2.show()

#resize, first image
image1 = image1.resize((128, 128))
image2= image2.resize((128, 128))

image1_size = image1.size
image2_size = image2.size

new_image = Image.new('RGB',(2*image1_size[0], image1_size[1]), (250,250,250))

new_image.paste(image1,(0,0))
new_image.paste(image2,(image1_size[0],0))

new_image.save("merged_image.png","PNG")
new_image.show()
