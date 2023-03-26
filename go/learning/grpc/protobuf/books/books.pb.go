// Just exploring about defining a proto file
// What data types supported? how are they defined? etc.
// Supports C style // and /* style */ comments

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: books/books.proto

package books

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// defining enums with enum keyword
type BookGenre int32

const (
	BookGenre_FICTION  BookGenre = 0
	BookGenre_THRILLER BookGenre = 1
	BookGenre_MEMOIR   BookGenre = 2
)

// Enum value maps for BookGenre.
var (
	BookGenre_name = map[int32]string{
		0: "FICTION",
		1: "THRILLER",
		2: "MEMOIR",
	}
	BookGenre_value = map[string]int32{
		"FICTION":  0,
		"THRILLER": 1,
		"MEMOIR":   2,
	}
)

func (x BookGenre) Enum() *BookGenre {
	p := new(BookGenre)
	*p = x
	return p
}

func (x BookGenre) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BookGenre) Descriptor() protoreflect.EnumDescriptor {
	return file_books_books_proto_enumTypes[0].Descriptor()
}

func (BookGenre) Type() protoreflect.EnumType {
	return &file_books_books_proto_enumTypes[0]
}

func (x BookGenre) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BookGenre.Descriptor instead.
func (BookGenre) EnumDescriptor() ([]byte, []int) {
	return file_books_books_proto_rawDescGZIP(), []int{0}
}

// Book attributes defined with singular scalar fields
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year       uint32    `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Price      float64   `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	IsReleased bool      `protobuf:"varint,4,opt,name=is_released,json=isReleased,proto3" json:"is_released,omitempty"`
	Genre      BookGenre `protobuf:"varint,5,opt,name=genre,proto3,enum=BookGenre" json:"genre,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_books_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_books_books_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Book) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Book) GetIsReleased() bool {
	if x != nil {
		return x.IsReleased
	}
	return false
}

func (x *Book) GetGenre() BookGenre {
	if x != nil {
		return x.Genre
	}
	return BookGenre_FICTION
}

// list defined using 'repeated'
type Bookshelf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *Bookshelf) Reset() {
	*x = Bookshelf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookshelf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookshelf) ProtoMessage() {}

func (x *Bookshelf) ProtoReflect() protoreflect.Message {
	mi := &file_books_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookshelf.ProtoReflect.Descriptor instead.
func (*Bookshelf) Descriptor() ([]byte, []int) {
	return file_books_books_proto_rawDescGZIP(), []int{1}
}

func (x *Bookshelf) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

// map defined using 'map'
type BooksByAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books map[string]*Bookshelf `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BooksByAuthor) Reset() {
	*x = BooksByAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksByAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksByAuthor) ProtoMessage() {}

func (x *BooksByAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_books_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksByAuthor.ProtoReflect.Descriptor instead.
func (*BooksByAuthor) Descriptor() ([]byte, []int) {
	return file_books_books_proto_rawDescGZIP(), []int{2}
}

func (x *BooksByAuthor) GetBooks() map[string]*Bookshelf {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_books_books_proto protoreflect.FileDescriptor

var file_books_books_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x73, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x20, 0x0a,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22,
	0x28, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x1b, 0x0a, 0x05,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x0d, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x1a, 0x44, 0x0a, 0x0a,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x2a, 0x32, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x46, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x54, 0x48, 0x52, 0x49, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45,
	0x4d, 0x4f, 0x49, 0x52, 0x10, 0x02, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x67, 0x79, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_books_proto_rawDescOnce sync.Once
	file_books_books_proto_rawDescData = file_books_books_proto_rawDesc
)

func file_books_books_proto_rawDescGZIP() []byte {
	file_books_books_proto_rawDescOnce.Do(func() {
		file_books_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_books_proto_rawDescData)
	})
	return file_books_books_proto_rawDescData
}

var file_books_books_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_books_books_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_books_books_proto_goTypes = []interface{}{
	(BookGenre)(0),        // 0: BookGenre
	(*Book)(nil),          // 1: Book
	(*Bookshelf)(nil),     // 2: Bookshelf
	(*BooksByAuthor)(nil), // 3: BooksByAuthor
	nil,                   // 4: BooksByAuthor.BooksEntry
}
var file_books_books_proto_depIdxs = []int32{
	0, // 0: Book.genre:type_name -> BookGenre
	1, // 1: Bookshelf.books:type_name -> Book
	4, // 2: BooksByAuthor.books:type_name -> BooksByAuthor.BooksEntry
	2, // 3: BooksByAuthor.BooksEntry.value:type_name -> Bookshelf
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_books_books_proto_init() }
func file_books_books_proto_init() {
	if File_books_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_books_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_books_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bookshelf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_books_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksByAuthor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_books_books_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_books_books_proto_goTypes,
		DependencyIndexes: file_books_books_proto_depIdxs,
		EnumInfos:         file_books_books_proto_enumTypes,
		MessageInfos:      file_books_books_proto_msgTypes,
	}.Build()
	File_books_books_proto = out.File
	file_books_books_proto_rawDesc = nil
	file_books_books_proto_goTypes = nil
	file_books_books_proto_depIdxs = nil
}
