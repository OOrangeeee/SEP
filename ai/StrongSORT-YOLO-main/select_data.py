import os
import shutil
dir = os.listdir('/memory/huqiang/Polyp/images/none_polyp/images')
save_path = '/memory/daiwenxuan/Polyp/AQS'
for dir_name in ['AQ']:
    with open('/home/daiwenxuan/code/Polyp/polyp_segmentation/StrongSORT-YOLO-main/data/'+dir_name+'.txt','r') as f:
        lines = f.readlines()
        for line in lines:
            num = line.strip()
            image_path = '/memory/huqiang/Polyp/images/none_polyp/images/'+dir_name+'/'+dir_name+'_'+str(num)+'.jpg'
            destination_path = os.path.join(save_path, dir_name+'_'+str(num)+'.jpg')
            shutil.copy(image_path, destination_path)