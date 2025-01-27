# MILESTONE 1: IELTS READING APP
## MÔN: LẬP TRÌNH WINDOWS 
## LỚP: 22_3 
## SV THỰC HIỆN: 
- Mai Nhật Nam - MSSV: 22120219
- Âu Lê Tuấn Nhật - MSSV: 22120250
- Nguyễn Thành Phát - MSSV: 22120263
## MỤC TIÊU NHÓM TRONG MILESTONE 1
- Với milestone này, nhóm chúng em đã đề ra mục tiêu là xây dựng giao diện sơ bộ cho app, bao gồm:
  - Trang đăng nhập/đăng ký
  - Trang chủ (Homepage)
  - Trang giới thiệu app (About Us)
  Với ước lượng công việc là khoảng 1/3 dự án cho cả 3 thành viên, mỗi bạn sẽ làm việc trong khoảng 3.3 giờ.

- Về chức năng đăng nhập (với việc tổ chức phân quyền (authentication) ở mức cơ bản), nhóm đã đề ra các mục tiêu cụ thể như sau:
  1. **Đăng nhập bằng tài khoản bình thường**: người dùng sẽ nhập email và mật khẩu để đăng nhập vào ứng dụng. Hỗ trợ thêm cho người dùng các chức năng như Remember Me và Forget Password
  2. **Đăng nhập bằng Tài khoản Google**: Người dùng khi click vào button "Sign in with Google" thì app sẽ hỗ trợ chuyển sang đăng nhập bằng Gmail
  3. **Đăng ký tài khoản**: Người dùng có thể đăng ký tài khoản bằng cách nhập tên, email, password bên tab "Register" và có 1 email xác minh danh tính người dùng được gửi đến.
  4. **Phân quyền tài khoản**: Người dùng khi đăng ký tài khoản bình thường sẽ tự động có role là "end_user" được lưu ở dưới database và chỉ được phép sử dụng những tab dành cho người dùng bình thường. Còn người dùng là quản trị viên khi đăng nhập sẽ có thêm 1 số chức năng và quyền truy cập đến các API khác trong server.
  5. **Giao diện MainWindow cho các chức năng đăng nhập/đăng ký**: Giao diện này sẽ hiển thị 1 số thứ như tab login đăng nhập bình thường + button hỗ trợ đăng nhập qua Google/ tab Register
  
- Về trang Homepage (trang chủ), nhóm đã đề ra trang Homepage này sẽ có những nội dung sau:
  1. **Thanh menu**: Ở thanh này sẽ hiển thị những tab sau: Logo app, Trang chủ, Reading&Listening, Full Test, Courses (chức năng thêm sau và đang được nhóm cân nhắc thêm), AboutUs, Avatar và tên người dùng.
  2. **Hiển thị thông tin cơ bản của App**: Phần giao diện này sẽ hiển thị những thông tin cơ bản nhất của app
  3. **Phần đề xuất cho người dùng**: Phần giao diện này sẽ hiển thị 1 số nội dung/bài đọc/ kiến thức hay để giới thiệu thêm cho người dùng
  4. **Lịch sử làm bài**: Phần giao diện này sẽ hiển thị lịch sử làm bài và sẽ hỗ trợ người dùng filter theo độ khó bài tập hay theo danh sách bài tập
  5. **Mục tiêu của bạn**: Phần giao diện này sẽ hỗ trợ người dùng đặt mục tiêu cá nhân của mình. (Ví dụ như mục tiêu điểm số, lịch thi)
  6. **Kế hoạch người dùng**: Người dùng có thể sử dụng thêm Calendar của app để cá nhân hóa lộ trình học tập của mình.
  7. **Footer của trang**: Chứa thông tin liên lạc và các thông tin liên quan đến bản quyền app

- Về trang About Us (trang giới thiệu app), nhóm đã đề ra trang này sẽ có những nội dung sau:
  1. **Thanh menu**: Ở thanh này sẽ hiển thị những tab sau: Logo app, Trang chủ, Reading&Listening, Full Test, Courses (chức năng thêm sau và đang được nhóm cân nhắc thêm), AboutUs, Avatar và tên người dùng.
  2. **Phần thông tin về app**: Phần này sẽ cung cấp thông tin về lợi ích của app muốn mang lại đến cho người dùng
  3. **Footer của trang**: Chứa thông tin liên lạc và các thông tin liên quan đến bản quyền app.

## CÁC CHỨC NĂNG MÀ NHÓM ĐÃ HOÀN THIỆN TRONG MILESTONE 1
- Với chức năng đăng nhập/đăng ký: App đã có thể đăng nhập qua email + password, remember me, đăng nhập qua google, đăng ký tài khoản thông thường.
- Với màn hình homepage: Hiện tại các mục đều đã hoàn thiện cơ bản nhưng vẫn còn 1 số lỗi được đề cập trong phần "CÁC CHỨC NĂNG CHƯA HOÀN THIỆN VÀ CÒN LỖI TRONG MILESTONE 1".
- Với màn hình AboutUs: Hiện tại các mục đều đã hoàn thiện cơ bản nhưng vẫn còn 1 số lỗi được đề cập trong phần "CÁC CHỨC NĂNG CHƯA HOÀN THIỆN VÀ CÒN LỖI TRONG MILESTONE 1".

## CÁC CHỨC NĂNG CHƯA HOÀN THIỆN VÀ CÒN LỖI TRONG MILESTONE 1
- Với màn hình đăng nhập/đăng ký: App của chúng em hiện tại chưa hỗ trợ forgot password và khi người dùng đăng ký tài khoản bình thường, nhóm chưa có email tự động xác nhận danh tính
- Với màn hình homepage: Phần hiển thị lịch sử làm bài hiện chưa có bài tập nào được ra do tính năng Luyện tập vẫn đang trong quá trình xây dựng và sẽ hoàn thiện ở milestone 2 (Transaction). Hơn nữa, phần đề xuất làm bài cũng với lý do tính năng luyện tập chưa có nên chưa hiển thị nội dung ra.
- Với cả 3 màn hình, nhóm chưa sắp xếp kịp để code phần Responsive và cũng như về cấu trúc source code vẫn còn đang lộn xộn. Chưa có nhiều các xử lí có bắt các lỗi cơ bản hay không (không có dữ liệu, dữ liệu ko đúng định dạng, dữ liệu ko đúng miền giá trị)
- Về vấn đề lưu trữ dữ liệu 
- Do đây là lần đầu nhóm em được tiếp xúc với quy trình phát triển 1 phần mềm nên tất nhiên nhóm sẽ chưa thể nào hoàn thiện được 1 cách trọn vẹn được và nhóm em cam kết sẽ cải thiện những điểm này ngay ở trong milestone tiếp theo.

## DESIGN PATTERNS/ ARCHITECTURE ĐƯỢC SỬ DỤNG TRONG MILESTONE 1
**Về Architecture**
- Với phần source FE: Nhóm em đang sử dụng mô hình MVVM nhưng chưa thật sự hiểu rõ cách cấu hình folder/files cho mô hình này nên sẽ có 1 kế hoạch điều chỉnh lớn về code structure trong milestone số 2. Hiện tại code structure của nhóm còn rất lộn xộn và chỉ đang cơ bản là tạo 1 màn hình .xaml và file cs của màn hình đó tương ứng. Và theo mô hình này thì nhóm em sẽ có kế hoạch chỉnh sửa define các file/Folder như sau:
```plaintext
login_full/
├── Models/
│   ├── UserProfile.cs
│   ├── UserTarget.cs
│   └── ... (Các lớp đại diện cho dữ liệu hoặc thực thể của ứng dụng)
├── ViewModels/
│   ├── HomePageViewModel.cs
│   ├── MainWindowViewModel.cs
│   └── ... (Các lớp ViewModel quản lý logic và dữ liệu cho từng View tương ứng)
├── Views/
│   ├── HomePage.xaml
│   ├── HomePage.xaml.cs
│   ├── MainWindow.xaml
│   ├── MainWindow.xaml.cs
│   └── ... (Các file XAML và code-behind đại diện cho giao diện người dùng)
├── Services/
│   ├── API/
│   │   ├── ConfigService.cs
│   │   ├── DatabaseService.cs
│   │   ├── GoogleAuthService.cs
│   │   ├── LoginApiService.cs
│   └── Managers/
│       ├── CalendarManager.cs
│       ├── ScheduleManager.cs
│       └── ... (Các lớp dịch vụ và quản lý logic nghiệp vụ không phụ thuộc vào giao diện)
├── Helpers/
│   └── ... (Các tiện ích hoặc hàm hỗ trợ có thể tái sử dụng trong nhiều phần khác nhau của ứng dụng)
├── Config/
│   ├── appsettings.json
│   ├── config.yaml
│   └── ... (Các file cấu hình cho ứng dụng)
├── App.xaml
├── App.xaml.cs
└── ... (Các file và thư mục khác nếu cần)
```
- Với phần source BE (ngôn ngữ Golang - framework GIN): Nhóm em đã xây dựng source dựa trên **Clean Architecture** với các thư mục được tổ chức rõ ràng, tách biệt trách nhiệm và phụ thuộc, giúp dễ bảo trì và mở rộng. Hơn nữa, với phần source BE nhóm tụi em đã deploy API lên server của Render. Tuy nhiên, Render sẽ cho hệ thống API vào trạng thái sleep nếu như sau 15' mà không có bất kì 1 API nào được request đến. Và nếu như server API bị sleep thì lần gọi API tiếp theo sẽ phải mất 1 khoảng thời gian khá lâu để gọi được API (>=50s) nên nhóm em đã có sử dụng thêm 1 service cron job free có tên là [cronjob.org] (https://cron-job.org/en/) để cứ mỗi 10', nhóm em sẽ cho con cron job này gọi đến 1 API Get bất kì để giữ server không bị sleep. 
  1. **cmd**: Thư mục này thường chứa các điểm khởi đầu chính của ứng dụng. Trong trường hợp này, có thể có một tệp chính trong `cmd` để khởi chạy ứng dụng.
  2. **common**: Thư mục này chứa các đoạn mã tiện ích và các chức năng dùng chung trong toàn bộ ứng dụng:
       - `db.go`: Chứa cấu hình và các hàm kết nối cơ sở dữ liệu.
       - `error.go`, `error_messages.go`: Xử lý các loại lỗi tùy chỉnh và thông điệp lỗi.
       - `helper.go`, `utils.go`: Các hàm trợ giúp chung có thể tái sử dụng trong các package khác.
       - `jwt.go`: Quản lý mã JWT, xử lý mã hóa và giải mã để xác thực.
  3. **config**: Thư mục này có thể chứa các thiết lập và cấu hình liên quan đến ứng dụng, như các tệp cấu hình hoặc các tham số cấu hình hệ thống.
  4. **internal**: Thư mục `internal` chứa các thành phần cốt lõi của ứng dụng, không được truy cập từ bên ngoài, và được chia thành các thư mục con theo vai trò:
     - handlers: chứa các tệp xử lý logic API như `authen.go` (Xử lý logic xác thực người dùng), `check_health.go` (Kiểm tra tình trạng của dịch vụ), `target.go`, `user.go`(Xử lý các chức năng liên quan đến đối tượng `target` và `user`)
     - models: chứa các tệp định nghĩa cấu trúc dữ liệu của ứng dụng (models) như `authen.go`, `user.go`, `target.go`, `role.go`, `base.go` để định nghĩa cấu trúc dữ liệu cho các đối tượng liên quan như `user`, `role` (vai trò), `target` (đối tượng), v.v.
     - repositories: chứa các tệp làm việc với cơ sở dữ liệu như `authen.go`, `user.go`, `target.go`, `base.go` để xử lý các thao tác CRUD (tạo, đọc, cập nhật, xóa) cho từng đối tượng cụ thể, đảm bảo tách biệt việc truy cập dữ liệu khỏi logic nghiệp vụ.
     - services: chứa các tệp xử lý logic nghiệp vụ như `authen.go`, `user.go`, `target.go`, `base.go` để xử lý logic nghiệp vụ cho từng đối tượng cụ thể, đóng vai trò trung gian giữa các `handler` và `repository`.
     - middleware: chứa các tệp xử lý trung gian (middleware) như  `admin_authentication.go`, `user_authentication.go` để quản lý xác thực và phân quyền, kiểm tra quyền truy cập của admin và user.
  5. **pkg**: để chứa các package bổ trợ như `postgres` (Chứa các tiện ích hoặc cấu hình dành riêng cho PostgreSQL) hay `utils` (Các tiện ích khác có thể cần thiết cho nhiều phần của ứng dụng).
  6. **script**: Chứa các script hoặc tệp SQL cần thiết cho ứng dụng
  7. **Các tệp cấu hình chung**:
  - `.env`: Tệp môi trường, chứa các biến môi trường như cấu hình cơ sở dữ liệu, thông tin API, v.v.
  - `go.mod`, `go.sum`: Quản lý các module và dependencies của Go.
  - `.gitignore`: Xác định các tệp và thư mục không cần thiết đẩy lên Git.
 
Cấu trúc này giúp phân chia rõ ràng giữa các tầng (layer) của ứng dụng:
- **handlers** cho việc xử lý HTTP.
- **services** cho logic nghiệp vụ.
- **repositories** cho truy cập dữ liệu.
- **models** cho các định nghĩa dữ liệu.



**Về design patterns** nhóm có sử dụng Singleton cho GlobalState để lưu trữ trạng thái ứng dụng để sử dụng GlobalState.Instance ở bất kỳ đâu trong ứng dụng để truy cập đến AccessToken hoặc UserProfile được lưu trữ trong GlobalState mà không cần phải khởi tạo lại lớp. Nhưng điểm nhóm em chưa kịp sửa lại đối với việc xử lý AccessToken đó là hash những thông tin nhạy cảm này đi ạ.

**Về database**:  
- Dưới đây là mô tả về các bảng trong cơ sở dữ liệu trong milestone 1 của nhóm: 

1. **Bảng `access_control_list`**
   Bảng này dùng để kiểm soát quyền truy cập của người dùng vào các hành động cụ thể trong hệ thống.
   - **Cột chính**:
     - `id`: Mã định danh duy nhất cho từng bản ghi.
     - `action_id`: Xác định hành động cụ thể mà người dùng có thể truy cập.
     - `role_id`: Liên kết với vai trò của người dùng để xác định quyền hạn.
     - `status`: Trạng thái của quyền truy cập.
     - `user_id`: Xác định người dùng cụ thể nếu cần.

2. **Bảng `roles`**
   Chứa thông tin về các vai trò trong hệ thống, mỗi vai trò có các quyền và hạn chế riêng.
   - **Cột chính**:
     - `id`: Mã định danh duy nhất cho vai trò.
     - `name`: Tên của vai trò (ví dụ: `Admin`, `End User`).
     - `enforce_tfa`: Cờ boolean cho biết vai trò này có yêu cầu xác thực hai yếu tố (TFA) hay không.
     - `admin_access`: Xác định quyền truy cập với tư cách admin.
     - `app_access`: Xác định quyền truy cập ứng dụng thông thường.
     - `public_access`: Quyền truy cập công khai (nếu có).

3. **Bảng `student_target`**
   Bảng này lưu các mục tiêu học tập của sinh viên cho các kỹ năng khác nhau.
   - **Cột chính**:
     - `id`: Mã định danh duy nhất cho từng sinh viên.
     - `target_study_duration`: Thời gian học tập mục tiêu của sinh viên.
     - `target_reading`: Điểm mục tiêu cho kỹ năng đọc.
     - `target_listening`: Điểm mục tiêu cho kỹ năng nghe.
     - `target_speaking`: Điểm mục tiêu cho kỹ năng nói.
     - `target_writing`: Điểm mục tiêu cho kỹ năng viết.

4. **Bảng `users`**
   Chứa thông tin về người dùng trong hệ thống.
   - **Các fiekd chính**:
     - `id`: Mã định danh duy nhất cho mỗi người dùng.
     - `first_name`: Tên của người dùng.
     - `last_name`: Họ của người dùng.
     - `email`: Email của người dùng.
     - `password`: Mật khẩu (có thể được mã hóa) của người dùng.
     - `location`: Vị trí của người dùng.
     - `title`: Chức danh của người dùng.
     - `tags`: Thẻ gán cho người dùng (JSON).
     - `avatar`: Mã định danh ảnh đại diện của người dùng.
     - `language`: Ngôn ngữ của người dùng.
     - `tfa_secret`: Bí mật dùng để thiết lập xác thực hai yếu tố.
     - `status`: Trạng thái tài khoản của người dùng.
     - `role`: Liên kết vai trò của người dùng (liên kết với bảng `roles`).
     - `token`: Thông tin token truy cập.
     - `last_access`: Thời gian truy cập cuối cùng.
     - `provider`: Nhà cung cấp dịch vụ xác thực.
     - `external_identifier`: Định danh bên ngoài của người dùng.
     - `auth_data`: Dữ liệu xác thực bổ sung.
     - `email_notifications`: Tùy chọn nhận thông báo qua email.
     - `birthday`: Ngày sinh của người dùng.
     - `referrer_email`: Email người giới thiệu.
     - `fullname`: Họ và tên đầy đủ.
     - `search`: Chỉ mục tìm kiếm.
     - `facebook_url`: URL Facebook của người dùng.
     - `verify_token`: Token xác thực.
     - `is_active`: Trạng thái hoạt động của tài khoản.
     - `phone_number`: Số điện thoại của người dùng.
---
Cấu trúc này giúp phân chia quyền hạn và chức năng của từng người dùng trong hệ thống thông qua các bảng `roles` và `access_control_list`. Mối quan hệ giữa các bảng cho phép xác định chi tiết quyền truy cập và các mục tiêu học tập của sinh viên.

## ADVANCED TOPIC CỦA NHÓM TRONG MILESTONE 1
- Các topic nâng cao mà nhóm đã sử dụng trong milestone này là
   1. **Chức năng đăng nhập sử dụng O-auth**: Nhóm có hỗ trợ người dùng sign in with Google mà không cần phải tạo tài khoản thông thường.
   2. **Code Server BE API**: Dù đây không phải là môn học phát triển Web nhưng nhóm em nhận thấy rằng trong quá trình code sẽ có những chỗ gọi đến API để hỗ trợ. Ví dụ như trong milestone này nhóm sẽ gọi API SignUp, Login, UserProfile, UserTarget để hỗ trợ trong việc đăng nhập/đăng ký, hiển thị thông tin cá nhân, thông tin về mục tiêu thi của người dùng. Với việc sử dụng API có sẵn (ví dụ như API của Postgres sẽ cần phải viết function sql thì mới chạy được) thì nhóm cảm thấy sau này sẽ có 1 số API cần logic rất nhiều (ví dụ như API Submit Quiz tức khi người dùng bấm submit quiz thì việc cần 1 API trả ra kết quả của người dùng với quiz đó là rất cần thiết) nên đó là lý do chính nhóm build 1 server API đã được deploy trên dịch vụ Render với ngôn ngữ server API nhóm sử dụng là Golang.
## TEAMWORK - GITFLOW TRONG MILESTONE 1
- Về phân chia Task làm cho các thành viên trong nhóm thầy có thểm xem ở link [Google Sheets này](https://docs.google.com/spreadsheets/d/1GFQCUE59nozbMPFGt0ETG0eUi8qhOQHoTp4rLa3z_YA/edit?usp=sharing). Nhóm sẽ thực hiện đồ án này theo mô hình Scrum và các task được chia theo sprint. 1 Sprint của nhóm được quy ước là 2 tuần và trong sheet trên là các task cho sprint 1 + 2. Và hết 2 Sprint này là cũng vừa lúc milestone 1 kết thúc
- Về Git: Nhóm sẽ có 2 source và được tạo trong 1 organization. 1 Source dành cho UI và Source dành cho BE API. 
  - Một số hình ảnh minh chứng về hoạt động Git của nhóm:
    - ![image](https://drive.google.com/uc?export=view&id=1Hmo95HaYzjyrkp2VDJoBtr9jbjBLpznr)
    - ![image 1](https://drive.google.com/uc?export=view&id=1lugRocpAOMUtUZQEOh8eBsN67gSja4Er)
    - ![image 3](https://drive.google.com/uc?export=view&id=1cbLzUi4JfDCGjNvrf0Riypu7TVjp8BO3)

## QUALITY ASSURANCE TRONG MILESTONE 1
- **Quá trình duyệt mã nguồn** Quy trình push code được nhóm quy định như sau:
  1. **main**: Ở branch này là nơi chứa source chính sau khi được kiểm thử thành công và nộp bài sau mỗi milestone.
  2. **feature/xx**: Các thành viên sẽ chủ yếu code trên branch này và branch này sẽ được checkout từ main để tạo ra
  3. **release/stg**: Sau khi các thành viên code xong, mọi người sẽ tạo pull requests vào branch này để review code. Sau khi mọi thứ được hoàn thiện thì branch này sẽ được lấy và merge vào main để "deploy".
  - Một số hình ảnh minh chứng về hoạt động kiểm duyệt mã nguồn của nhóm:
    - ![image 1](https://drive.google.com/uc?export=view&id=1B0kQQ52Q31pburOZooq93FOzB6QiBdGq)
    - ![image 2](https://drive.google.com/uc?export=view&id=1m6yCYhSj3FyE3RJsL2x4vJ4_fGO2SNlR)
    - ![image 3](https://drive.google.com/uc?export=view&id=1l5atm3mEMPYI9sZDCUrIGAtFV1p9Voz-)
- **UI Test**: Ngoài ra nhóm kết hợp với việc viết UI test cũng như là mannual testing để đảm bảo ít lỗi xảy ra nhất. Các test case thầy có thể tìm thấy ở link [Google Sheets này](https://docs.google.com/spreadsheets/d/1tb8dCWPa4k6dqzmSZ1OEZAzH4eXli-zu/edit?usp=sharing&ouid=101425335323710410200&rtpof=true&sd=true) 

## HƯỚNG DẪN GIÁO VIÊN CHẤM BÀI CỦA NHÓM TRONG MILESTONE 1
- Do app chưa được deploy và còn chạy ở local nên khi đăng nhập qua Oauth thì để API `https://oauth2.googleapis.com/tokeninfo?id_token` trả ra field `name` để phía API Login sẽ sử dụng field `name` này để ghi vào database nên nhóm em cần người dùng sẽ phải thao tác thiết lập OAuth trên Google Console đúng cách để ứng dụng có quyền truy cập các thông tin cần thiết (chẳng hạn như email và user ID của người dùng).
- Các bước thao tác như sau:

  1. **Tạo OAuth Consent Screen (Màn hình đồng ý OAuth)**
      - Vào [Google Cloud Console](https://console.cloud.google.com/).
      - Chọn dự án hoặc tạo dự án mới.
      - Điều hướng tới **APIs & Services > OAuth consent screen**.
      - Cấu hình các thông tin cơ bản cho màn hình đồng ý, chẳng hạn như tên ứng dụng, địa chỉ email liên hệ, v.v.
      - Thêm các **scope** bạn muốn (như email và user ID). Những scope này sẽ yêu cầu người dùng cho phép chia sẻ các thông tin cần thiết khi họ đăng nhập.
      - Bấm Update để đồng ý cập nhật.

  2. **Tạo OAuth 2.0 Client ID**
      - Điều hướng tới **APIs & Services > Credentials**.
      - Nhấn vào **Create Credentials** và chọn **OAuth client ID**.
      - Trong mục **Application type**, chọn **Web application**.
      - Điền các thông tin cần thiết:
      - Tạo client ID và client secret cho ứng dụng của bạn.

  3. **Cấu hình scope cho email và user ID**
      - Trong Google Console, trong quá trình thiết lập OAuth, hãy thêm các scope sau:
        - `https://www.googleapis.com/auth/userinfo.email` để truy cập email của người dùng.
        - `https://www.googleapis.com/auth/userinfo.profile` để truy cập thông tin cơ bản của người dùng như tên và ID.
      - Điều này sẽ đảm bảo rằng khi người dùng đăng nhập qua OAuth, ứng dụng sẽ có quyền truy cập thông tin email và ID của họ.


## TỔNG KẾT - ĐÁNH GIÁ CỦA NHÓM TRONG MILESTONE 1
- Nhìn chung nhóm đã hoàn thành được 80% mục tiêu đề ra tuy vẫn còn rất nhiều thiếu sót lớn nên nhóm em đã rút kinh nghiệm và sẽ tập trung sửa các lỗi còn được nêu ở trên đối với milestone 2






