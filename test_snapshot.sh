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
spawn ./PackageDelivery


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
# expect -re "Masukkan pilihan \\(1/2/0\\):"
# type_input_no_enter "2"
# screenshot "1.select_menu" "flow/user"
# send -- "\r"

# # 2. Type username as user role
# expect -re "Masukkan username:"
# type_input_no_enter "testuser"
# screenshot "2.type_user_username" "flow/user"
# send -- "\r"

# # 3. Type password as user role
# expect -re "Masukkan password:"
# type_input_no_enter "testpass"
# screenshot "3.type_user_password" "flow/user"
# send -- "\r"




# # 1. Select “add package”
# expect -re "Pilih menu:"
# type_input_no_enter "1"
# screenshot "1.select_user_menu" "flow/user/add_package"
# send -- "\r"

# # 2. Type package type
# expect -re "Masukkan tipe paket:"
# type_input_no_enter "Reguler"
# screenshot "2.type_package_type" "flow/user/add_package"
# send -- "\r"

# # 3. Type package weight
# expect -re "Masukkan berat paket \\(kg\\):"
# type_input_no_enter "5"
# screenshot "3.type_package_weight" "flow/user/add_package"
# send -- "\r"

# # 4. Type package sender
# expect -re "Masukkan kota pengirim:"
# type_input_no_enter "Jakarta"
# screenshot "4.type_package_sender" "flow/user/add_package"
# send -- "\r"

# # 5. Type package receiver
# expect -re "Masukkan kota tujuan:"
# type_input_no_enter "Bekasi"
# screenshot "5.type_package_receiver" "flow/user/add_package"
# send -- "\r"

# expect -re "Jarak antar kota: 16.33 km"
# expect -re "Harga paket: Rp 81651"

# # 6. Confirm package
# expect -re "Apakah Anda yakin ingin menambahkan paket ini\\? \\(y/n\\):"
# type_input_no_enter "y"
# screenshot "6.confirm_package" "flow/user/add_package"
# send -- "\r"

# # 7. Package added successfully
# expect -re "Paket berhasil ditambahkan."
# screenshot "7.package_added" "flow/user/add_package"
# send -- "\r"




# # 1. Select “Cek Paket”
# expect -re "Pilih menu:"
# type_input_no_enter "2"
# screenshot "1.select_user_menu" "flow/user/check_package"
# send -- "\r"

# # 2. Type tracking number
# expect -re "Masukkan No Resi:"
# type_input_no_enter "6"
# screenshot "2.type_tracking_number" "flow/user/check_package"
# send -- "\r"

# # 3. Check package status
# expect -re "No Resi: 6"
# expect -re "Tipe: Reguler"
# expect -re "Berat: 5"
# expect -re "Harga: 81651"
# expect -re "Kota Pengirim: Jakarta"
# expect -re "Kota Tujuan: Bekasi"
# expect -re {Status: \[Paket Dibuat\]}
# expect -re "Dibuat pada:"
# expect -re "Diperbarui pada:"
# expect -re "Kurir: Belum Ditugaskan"
# expect -re "Paket berhasil ditambahkan."
# screenshot "3.check_package_status" "flow/user/check_package"
# send -- "\r"




# # 4. Logout To Continue with Admin Role
# expect -re "Pilih menu:"
# type_input_no_enter "0"
# screenshot "4.logout_as_user" "flow/user"
# send -- "\r"

# # ------------------------------------------------------------------------------



# # ------------------------------------------------------------------------------
# # role admin flow
# # ------------------------------------------------------------------------------
# # 1. Select “login”
# expect -re "Masukkan pilihan \\(1/2/0\\):"
# type_input_no_enter "2"
# screenshot "1.select_menu" "flow/admin"
# send -- "\r"

# # 2. Type username as admin role
# expect -re "Masukkan username:"
# type_input_no_enter "admin"
# screenshot "2.type_admin_username" "flow/admin"
# send -- "\r"

# # 3. Type password as admin role
# expect -re "Masukkan password:"
# type_input_no_enter "1"
# screenshot "2.2.3.type_admin_password" "flow/admin"
# send -- "\r"




# # 1. Select “add package”
# expect -re "Pilih menu:"
# type_input_no_enter "1"
# screenshot "1.select_user_menu" "flow/admin/add_package"
# send -- "\r"

# # 2. Type package type
# expect -re "Masukkan tipe paket:"
# type_input_no_enter "Express"
# screenshot "2.type_package_type" "flow/admin/add_package"
# send -- "\r"

# # 3. Type package weight
# expect -re "Masukkan berat paket \\(kg\\):"
# type_input_no_enter "10"
# screenshot "3.type_package_weight" "flow/admin/add_package"
# send -- "\r"

# # 4. Type package sender
# expect -re "Masukkan kota pengirim:"
# type_input_no_enter "Depok"
# screenshot "4.type_package_sender" "flow/admin/add_package"
# send -- "\r"

# # 5. Type package receiver
# expect -re "Masukkan kota tujuan:"
# type_input_no_enter "Bekasi"
# screenshot "5.type_package_receiver" "flow/admin/add_package"
# send -- "\r"

# expect -re "Jarak antar kota: 26.01 km"
# expect -re "Harga paket: Rp 130041"

# # 6. Confirm package
# expect -re "Apakah Anda yakin ingin menambahkan paket ini\\? \\(y/n\\):"
# type_input_no_enter "y"
# screenshot "6.confirm_package" "flow/admin/add_package"
# send -- "\r"

# # 7. Package added successfully
# expect -re "Paket berhasil ditambahkan."
# screenshot "7.package_added" "flow/admin/add_package"
# send -- "\r"


# # 1. Select "lihat semua paket"
# expect -re "Pilih menu:"
# type_input_no_enter "2"
# screenshot "1.select_view_all_packages" "flow/admin/view_all_packages"
# send -- "\r"

# # 2. expect added packages
# expect -re "6\\s+Express\\s+10\\.0\\s+130041\\s+Depok\\s+Bekasi"
# screenshot "1.view_all_packages" "flow/admin/view_all_packages"
# send -- "\r"




# # 1. Select "edit paket"
# expect -re "Pilih menu:"
# type_input_no_enter "3"
# screenshot "1.select_edit_package" "flow/admin/edit_package"
# send -- "\r"

# # 2. Type tracking number
# expect -re "Masukkan No Resi Paket yang ingin diedit:"
# type_input_no_enter "7"
# screenshot "2.type_tracking_number" "flow/admin/edit_package"
# send -- "\r"

# # 3. Type new package type
# expect -re {Tipe \[Express\]:}
# type_input_no_enter "Reguler"
# screenshot "3.type_new_package_type" "flow/admin/edit_package"
# send -- "\r"

# # 4. Type new package weight
# expect -re {Berat \[10\\.00\\]:}
# type_input_no_enter "5"
# screenshot "4.type_new_package_weight" "flow/admin/edit_package"
# send -- "\r"

# # 5. Type price
# expect -re {Harga \[130041\\]:}
# type_input_no_enter "65020"
# screenshot "5.type_new_package_price" "flow/admin/edit_package"
# send -- "\r"

# # 6. Type new sender city
# expect -re {Kota Pengirim \[Depok\\]:}
# type_input_no_enter "Jakarta"
# screenshot "6.type_new_package_sender" "flow/admin/edit_package"
# send -- "\r"

# # 7. Type new receiver city
# expect -re {Kota Tujuan \[Bekasi\\]:}
# type_input_no_enter "Bogor"
# screenshot "7.type_new_package_receiver" "flow/admin/edit_package"
# send -- "\r"

# # 8. Update package status
# expect -re {Status \(pisah dengan koma jika banyak\) \[Paket Dibuat\]:}
# type_input_no_enter "Paket Dibuat, Paket Ditolak"
# screenshot "8.update_package_status" "flow/admin/edit_package"
# send -- "\r"

# # 9. CHECK PACKAGE STATUS
# expect -re "No Resi: 6"
# expect -re "Tipe: Reguler"
# expect -re "Berat: 5.00"
# expect -re "Harga: 65020.00"
# expect -re "Kota Pengirim: Jakarta"
# expect -re "Kota Tujuan: Bogor"
# expect -re {Status: \[Paket Dibuat  Paket Ditolak\]}
# screenshot "9.check_updated_package_status" "flow/admin/edit_package"
# send -- "\r"




# # 1. Select "hapus paket"
# expect -re "Pilih menu:"
# type_input_no_enter "4"
# screenshot "1.select_delete_package" "flow/admin/delete_package"
# send -- "\r"

# # 2. Type tracking number
# expect -re "Masukkan No Resi Paket yang ingin dihapus:"
# type_input_no_enter "6"
# screenshot "2.type_tracking_number" "flow/admin/delete_package"
# send -- "\r"

# # 3. expect success message
# expect -re "Paket dengan No Resi 6 telah dihapus."
# screenshot "3.delete_package_success" "flow/admin/delete_package"
# send -- "\r"




# # 1. Select "tambah kurir baru"
# expect -re "Pilih menu:"
# type_input_no_enter "5"
# screenshot "1.select_add_courier" "flow/admin/add_courier"
# send -- "\r"

# # 2. Type courier name
# expect -re "Masukkan Nama:"
# type_input_no_enter "adam"
# screenshot "2.type_courier_name" "flow/admin/add_courier"
# send -- "\r"

# # 3. Type courier username
# expect -re "Masukkan Username:"
# type_input_no_enter "adamst"
# screenshot "3.type_courier_username" "flow/admin/add_courier"
# send -- "\r"

# # 4. Type courier password
# expect -re "Masukkan Password:"
# type_input_no_enter "adam123"
# screenshot "4.type_courier_password" "flow/admin/add_courier"
# send -- "\r"

# # 5. expect success message
# expect -re "Kurir berhasil ditambahkan."
# screenshot "5.add_courier_success" "flow/admin/add_courier"
# send -- "\r"




# # 1. Select "lihat semua kurir"
# expect -re "Pilih menu:"
# type_input_no_enter "6"
# screenshot "1.view_all_couriers" "flow/admin/view_all_couriers"
# send -- "\r"

# # 2. expect all couriers
# expect -re "adam\\s+adamst\\s+adam123\\s+0"
# expect -re "kurir2\\s+kurir2\\s+1\\s+1"
# expect -re "kurir\\s+kurir\\s+1\\s+2"
# screenshot "2.check_all_couriers" "flow/admin/view_all_couriers"
# send -- "\r"




# # 1. Select "edit kurir"
# expect -re "Pilih menu:"
# type_input_no_enter "7"
# screenshot "1.select_edit_courier" "flow/admin/edit_courier"
# send -- "\r"

# # 2. Type courier username
# expect -re "Masukkan Username Kurir yang ingin diedit:"
# type_input_no_enter "kurir"
# screenshot "2.type_courier_username_to_edit" "flow/admin/edit_courier"
# send -- "\r"

# # 3. Type new courier name
# expect -re {Nama \[kurir\]:}
# type_input_no_enter "doe"
# screenshot "3.type_new_courier_name" "flow/admin/edit_courier"
# send -- "\r"

# # 4. Type new courier username
# expect -re {Username \[kurir\]:}
# type_input_no_enter "doe"
# screenshot "4.type_new_courier_username" "flow/admin/edit_courier"
# send -- "\r"

# # 5. Type new courier password
# expect -re {Password \[1\]:}
# type_input_no_enter "doe123"
# screenshot "5.type_new_courier_password" "flow/admin/edit_courier"
# send -- "\r"

# # 6. Set new role
# expect -re {Role \[kurir\]:}
# type_input_no_enter "kurir"
# screenshot "6.set_new_courier_role" "flow/admin/edit_courier"
# send -- "\r"

# # 7. expect success message
# expect -re "Kurir berhasil diubah:"
# screenshot "7.edit_courier_success" "flow/admin/edit_courier"
# send -- "\r"




# # 1. Select "lihat semua kurir"
# expect -re "Pilih menu:"
# type_input_no_enter "6"
# screenshot "1.view_all_couriers1" "flow/admin/view_all_couriers"
# send -- "\r"

# # 2. expect all couriers
# expect -re "adam\\s+adamst\\s+adam123\\s+0"
# expect -re "kurir2\\s+kurir2\\s+1\\s+1"
# expect -re "doe\\s+doe\\s+doe123\\s+2"
# screenshot "2.check_all_couriers1" "flow/admin/view_all_couriers"
# send -- "\r"




# # 1.Select "hapus kurir"
# expect -re "Pilih menu:"
# type_input_no_enter "8"
# screenshot "1.select_delete_courier" "flow/admin/delete_courier"
# send -- "\r"

# # 2. Type courier username
# expect -re "Masukkan username kurir yang ingin dihapus:"
# type_input_no_enter "doe"
# screenshot "2.type_courier_username_to_delete" "flow/admin/delete_courier"
# send -- "\r"

# # 3. expect success message
# expect -re "Kurir dengan username doe telah dihapus."
# screenshot "3.delete_courier_success" "flow/admin/delete_courier"
# send -- "\r"




# # 1. Select "lihat semua kurir"
# expect -re "Pilih menu:"
# type_input_no_enter "6"
# screenshot "1.view_all_couriers2" "flow/admin/view_all_couriers"
# send -- "\r"

# # 2. expect all couriers
# expect -re "adam\\s+adamst\\s+adam123\\s+0"
# expect -re "kurir2\\s+kurir2\\s+1\\s+1"
# screenshot "2.check_all_couriers2" "flow/admin/view_all_couriers"
# send -- "\r"




# # 1. Select "Assign Paket Ke Kurir"
# expect -re "Pilih menu:"
# type_input_no_enter "9"
# screenshot "1.select_assign_package_to_courier" "flow/admin/assign_package_to_courier"
# send -- "\r"

# # 2. Type tracking number
# expect -re "Masukkan No Resi Paket yang ingin diassign:"
# type_input_no_enter "4"
# screenshot "2.type_tracking_number_to_assign" "flow/admin/assign_package_to_courier"
# send -- "\r"

# # 3. Type courier username
# expect -re "Masukkan Username Kurir yang akan mengantarkan paket:"
# type_input_no_enter "kurir2"
# screenshot "3.type_courier_username_to_assign" "flow/admin/assign_package_to_courier"
# send -- "\r"

# # 4. expect success message
# expect -re "Paket berhasil diassign ke kurir."
# screenshot "4.assign_package_success" "flow/admin/assign_package_to_courier"
# send -- "\r"


# # 1. Select "lihat semua paket"
# expect -re "Pilih menu:"
# type_input_no_enter "2"
# screenshot "1.view_all_packages1" "flow/admin/view_all_packages"
# send -- "\r"

# # 2. expect all packages
# expect -re "4\\s+Reguler\\s+4\\.0\\s+50000\\s+Jakarta\\s+Bekasi\\s+Paket Dibuat\\s+\\d\\{4\\}-\\d\\{2\\}-\\d\\{2\\}\\s\\d\\{2\\}:\\d\\{2\\}:\\d\\{2\\}\\s+\\d\\{4\\}-\\d\\{2\\}-\\d\\{2\\}\\s\\d\\{2\\}:\\d\\{2\\}:\\d\\{2\\}\\s+kurir2"
# screenshot "2.check_all_packages1" "flow/admin/view_all_packages"
# send -- "\r"

# # 4. Logout To Continue with Kurir Role
# expect -re "Pilih menu:"
# type_input_no_enter "0"
# screenshot "4.logout_as_admin" "flow/admin"
# send -- "\r"

# # ------------------------------------------------------------------------------

# # 2.3. Login flow as kurir
# # 2.3.1 Select “login”
# expect -re "Masukkan pilihan \\(1/2/0\\):"
# type_input_no_enter "2"
# screenshot "2.3.1.select_menu" "flow/kurir"
# send -- "\r"

# # 2.3.2 Type username as user role
# expect -re "Masukkan username:"
# type_input_no_enter "kurir2"
# screenshot "2.3.2.type_kurir_username" "flow/kurir"
# send -- "\r"

# # 2.3.3. Type password as user role
# expect -re "Masukkan password:"
# type_input_no_enter "1"
# screenshot "2.3.3.type_kurir_password" "flow/kurir"
# send -- "\r"




# # 1. Select “lihat paket yang diassign”
# expect -re "Pilih menu:"
# type_input_no_enter "1"
# screenshot "1.select_view_assigned_packages" "flow/kurir/view_assigned_packages"
# send -- "\r"

# # 2. expect assigned packages
# # some expect logic
# screenshot "2.check_assigned_packages" "flow/kurir/view_assigned_packages"
# send -- "\r"




# # 1. Select “update status paket”
# expect -re "Pilih menu:"
# type_input_no_enter "2"
# screenshot "1.select_update_package_status" "flow/kurir/update_package_status"
# send -- "\r"

# # 2. Type tracking number
# expect -re "Masukkan nomor urut paket yang ingin diupdate statusnya:"
# type_input_no_enter "2"
# screenshot "2.type_tracking_number_to_update" "flow/kurir/update_package_status"
# send -- "\r"

# # 3. Type new status
# expect -re "Masukkan status baru paket (contoh: Diambil, Dalam Pengiriman, Terkirim):"
# type_input_no_enter "Terkirim"
# screenshot "3.type_new_package_status" "flow/kurir/update_package_status"
# send -- "\r"

# # 4. expect success message
# expect -re "Status paket berhasil diperbarui."
# screenshot "4.update_package_status_success" "flow/kurir/update_package_status"
# send -- "\r"

# # 2.3.4. Logout To Continue with flow
# expect -re "Pilih menu:"
# type_input_no_enter "0"
# screenshot "2.3.4.logout_as_kurir" "flow/kurir"
# send -- "\r"

# # ------------------------------------------------------------------------------

