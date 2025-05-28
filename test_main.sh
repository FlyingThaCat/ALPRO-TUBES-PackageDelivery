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
    # foreach ch [split $input ""] {
        send -- $input
    # }
    expect -timeout 1 eof
}

# ------------------------------------------------------------------------------
# Main flow
# ------------------------------------------------------------------------------
spawn ./PackageDelivery-NoInject


# ------------------------------------------------------------------------------
# 1. Select “register”
expect -re "Masukkan pilihan \\(1/2/0\\):"
type_input_no_enter "1"
screenshot "1.select_menu" "registerAsUser"
send -- "\r"

# 2. Type username
expect -re "Masukkan username:"
type_input_no_enter "testuser"
screenshot "2.input_username" "registerAsUser"
send -- "\r"

# 3. Type password
expect -re "Masukkan password:"
type_input_no_enter "testpass"
screenshot "3.input_password" "registerAsUser"
send -- "\r"
# ------------------------------------------------------------------------------

# ------------------------------------------------------------------------------
# role user flow
# ------------------------------------------------------------------------------
# 1. Select “login”
expect -re "Masukkan pilihan \\(1/2/0\\):"
type_input_no_enter "2"
screenshot "1.select_menu" "flow/user"
send -- "\r"

# 2. Type username as user role
expect -re "Masukkan username:"
type_input_no_enter "testuser"
screenshot "2.type_user_username" "flow/user"
send -- "\r"

# 3. Type password as user role
expect -re "Masukkan password:"
type_input_no_enter "testpass"
screenshot "3.type_user_password" "flow/user"
send -- "\r"




# 1. Select “add package”
expect -re "Pilih menu:"
type_input_no_enter "1"
screenshot "1.select_user_menu" "flow/user/add_package"
send -- "\r"

# 2. Type package type
expect -re "Masukkan tipe paket:"
type_input_no_enter "Reguler"
screenshot "2.type_package_type" "flow/user/add_package"
send -- "\r"

# 3. Type package weight
expect -re "Masukkan berat paket \\(kg\\):"
type_input_no_enter "5"
screenshot "3.type_package_weight" "flow/user/add_package"
send -- "\r"

# 4. Type package sender
expect -re "Masukkan kota pengirim:"
type_input_no_enter "Jakarta"
screenshot "4.type_package_sender" "flow/user/add_package"
send -- "\r"

# 5. Type package receiver
expect -re "Masukkan kota tujuan:"
type_input_no_enter "Bekasi"
screenshot "5.type_package_receiver" "flow/user/add_package"
send -- "\r"

expect -re "Jarak antar kota: 16.33 km"
expect -re "Harga paket: Rp 81651"

# 6. Confirm package
expect -re "Apakah Anda yakin ingin menambahkan paket ini? \\(y/n\\):"
type_input_no_enter "y"
screenshot "6.confirm_package" "flow/user/add_package"
send -- "\r"

# 7. Package added successfully
expect -re "Paket berhasil ditambahkan."
screenshot "7.package_added" "flow/user/add_package"
send -- "\r"




# 1. Select “Cek Paket”
expect -re "Pilih menu:"
type_input_no_enter "2"
screenshot "1.select_user_menu" "flow/user/check_package"
send -- "\r"

# 2. Type tracking number
expect -re "Masukkan No Resi:"
type_input_no_enter "1"
screenshot "2.type_tracking_number" "flow/user/check_package"
send -- "\r"

# 3. Check package status
expect -re "No Resi: 1"
expect -re "Tipe: Reguler"
expect -re "Berat: 1"
expect -re "Harga: 10000"
expect -re "Kota Pengirim: Jakarta"
expect -re "Kota Tujuan: Depok"
expect -re {Status: \[Paket Dibuat\]}
expect -re "Dibuat pada:"
expect -re "Diperbarui pada:"
expect -re "Kurir: kurir"
expect -re "Paket berhasil ditambahkan."
screenshot "3.check_package_status" "flow/user/check_package"
send -- "\r"




# 4. Logout To Continue with Admin Role
expect -re "Pilih menu:"
type_input_no_enter "0"
screenshot "4.logout_as_user" "flow/user"
send -- "\r"

# ------------------------------------------------------------------------------

# 2.2. Login flow as admin
# 2.2.1 Select “login”
expect -re "Masukkan pilihan \\(1/2/0\\):"
type_input_no_enter "2"
screenshot "2.2.1.select_menu" "login"
send -- "\r"

# 2.2.2 Type username as user role
expect -re "Masukkan username:"
type_input_no_enter "admin"
screenshot "2.2.2.type_admin_username" "login"
send -- "\r"

# 2.2.3. Type password as user role
expect -re "Masukkan password:"
type_input_no_enter "1"
screenshot "2.2.3.type_admin_password" "login"
send -- "\r"

# 2.2.4. Logout To Continue with Admin Role
expect -re "Pilih menu:"
type_input_no_enter "0"
screenshot "2.2.4.logout_as_admin" "login"
send -- "\r"

# ------------------------------------------------------------------------------

# 2.3. Login flow as kurir
# 2.3.1 Select “login”
expect -re "Masukkan pilihan \\(1/2/0\\):"
type_input_no_enter "2"
screenshot "2.3.1.select_menu" "login"
send -- "\r"

# 2.3.2 Type username as user role
expect -re "Masukkan username:"
type_input_no_enter "kurir"
screenshot "2.3.2.type_kurir_username" "login"
send -- "\r"

# 2.3.3. Type password as user role
expect -re "Masukkan password:"
type_input_no_enter "1"
screenshot "2.3.3.type_kurir_password" "login"
send -- "\r"

# 2.3.4. Logout To Continue with flow
expect -re "Pilih menu:"
type_input_no_enter "0"
screenshot "2.3.4.logout_as_kurir" "login"
send -- "\r"

# ------------------------------------------------------------------------------



# ------------------------------------------------------------------------------

# # --- Logout flow (if you still need it) ---
# expect -re "Masukkan pilihan \\(1/2/0\\):"
# type_input_no_enter "0" 0.05
# screenshot "select_menu" "logout"
# send -- "\r"
# sleep 0.5

# expect eof
