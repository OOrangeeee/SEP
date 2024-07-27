# path = '/memory/huqiang/Polyp/images/none_polyp/images'
# for dir in os
dir="/memory/huqiang/Polyp/images/none_polyp/images/"
for name in "CY" "CYC" "DXM" "FCF" "FWQ" "GS" "GZY" "HB" "HMF" "HMQ" "HQY" "HX" "HZX" "LC" "LCH" "LHY" "LJL" "LPB" "LQ" "LYZ" "MHT" "ML" "MTL" "MYH" "MZQ" "PJX" "TS" "WB" "WDB" "WGX" "WLJ" "WXH" "XSB" "XYL" "XYX" "YOUJY" "ZCH" "ZCL" "ZCP" "ZSX" "ZXY" "ZY"
do
    python /home/daiwenxuan/code/StrongSORT-YOLO-main/track_v5.py --source "$dir$name"
done
