#!/usr/bin/expect -f

# install expect, xdotool, and imagemagick if not already installed
# sudo apt-get install expect xdotool imagemagick
# run it on new terminal with:
# ./test.sh
# DONT INTERACT WITH ANYTHING ELSE WHILE THIS IS RUNNING

# ------------------------------------------------------------------------------
# Setup
# ------------------------------------------------------------------------------
exec mkdir -p logs/

proc ensure_dir {dir} {
    if {![file exists $dir]} {
        file mkdir $dir
    }
}

proc screenshot {name subfolder} {
    set folder "logs/$subfolder"
    ensure_dir $folder
    set filename "$folder/$name.png"
    exec sh -c "import -window \$(xdotool getactivewindow) '$filename'"
}

proc type_input_no_enter {input} {
    foreach ch [split $input ""] {
        send -- $ch
    }
    expect -timeout 1 eof
}

# ------------------------------------------------------------------------------
# Main flow
# ------------------------------------------------------------------------------
spawn ./PackageDelivery

# --- Register flow (unchanged) ---
expect -re "Masukkan pilihan \\(1/2/0\\):"
type_input_no_enter "1"
screenshot "1.select_menu" "register"
send -- "\r"

expect -re "Masukkan username:"
type_input_no_enter "testuser"
screenshot "2.input_username" "register"
send -- "\r"

expect -re "Masukkan password:"
type_input_no_enter "testpass"
screenshot "3.input_password" "register"
send -- "\r"

# # --- Login flow (explicit steps for two screenshots) ---
# # 1) Select “login”
# expect -re "Masukkan pilihan \\(1/2/0\\):"
# type_input_no_enter "2" 0.05
# screenshot "select_menu" "login"
# send -- "\r"
# sleep 0.5

# # 2) Type username → screenshot → Enter
# expect -re "Masukkan username:"
# type_input_no_enter "testuser" 0.05
# exec sleep 0.2
# screenshot "login_username" "login"
# exec sleep 0.2
# send -- "\r"
# sleep 0.5

# # 3) Type password → screenshot (now the username line is still on the screen) → Enter
# expect -re "Masukkan password:"
# type_input_no_enter "testpass" 0.05
# exec sleep 0.2
# screenshot "login_password" "login"
# exec sleep 0.2
# send -- "\r"
# sleep 0.5

# # --- Logout flow (if you still need it) ---
# expect -re "Masukkan pilihan \\(1/2/0\\):"
# type_input_no_enter "0" 0.05
# screenshot "select_menu" "logout"
# send -- "\r"
# sleep 0.5

# expect eof
