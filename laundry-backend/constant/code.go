package constant

const (
	SuccessCode       = "00"
	CreatedCode       = "01"
	UpdatedCode       = "02"
	DeletedCode       = "03"
	NotFoundCode      = "04"
	BadRequestCode    = "05"
	ConflictCode      = "06"
	UnprocessableCode = "07"
	InternalErrorCode = "99"

	SuccessMessage       = "OK	Sukses memproses permintaan"
	CreatedMessage       = "Created	Data berhasil ditambahkan"
	UpdatedMessage       = "OK	Data berhasil diperbarui"
	DeletedMessage       = "OK	Data berhasil dihapus"
	NotFoundMessage      = "Not Found	Data tidak ditemukan (misalnya ID parent/child tidak valid)"
	BadRequestMessage    = "Bad Request	Validasi gagal (misalnya field kosong, format salah)"
	ConflictMessage      = "Conflict	Terjadi konflik hirarki (misalnya parent sudah punya child unik tertentu)"
	UnprocessableMessage = "Unprocessable Entity	Hirarki tidak valid (misalnya loop/circular reference)"
	InternalErrorMessage = "Internal Server Error	Error tak terduga pada server"
)
const (
	RcSuccess          = "00"
	RmSuccess          = "Success"
	RcFailed           = "ADMIN"
	RmFailed           = "ADMIN"
	RcInvalidParameter = "ADMIN"
	RcNotFound         = "ADMIN"
)
