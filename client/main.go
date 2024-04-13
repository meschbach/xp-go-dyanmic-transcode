package main

import (
	"context"
	"fmt"
	"github.com/meschbach/xp-go-dyanmic-transcode/ipc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9432", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	v := &ipc.OtherMessage{
		Str: "helpful",
		I:   16,
	}
	wire, err := anypb.New(v)
	if err != nil {
		panic(err)
	}
	//todo: figure out why this is prepended -- turns out it is a sanity check by the protobuf library.  not sure why
	// this sanity check exists, so I am just stripping it.
	wire.TypeUrl = wire.TypeUrl[len("type.googleapis.com/"):]
	fmt.Printf("Wire: %s\n", wire.TypeUrl)
	format := FileDescriptorSetForMessage(v.ProtoReflect().Descriptor())

	client := ipc.NewTranscodeClient(conn)
	out, err := client.ToJSON(context.Background(), &ipc.JsonIn{
		What:   wire,
		Format: format,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(out.Json)
}

// FileDescriptorSetForMessage returns the FileDescriptorSet needed to represent m.
// copied (and refed message) from https://github.com/golang/protobuf/issues/1215#issuecomment-713192393
func FileDescriptorSetForMessage(m protoreflect.Descriptor) *descriptorpb.FileDescriptorSet {
	fds := new(descriptorpb.FileDescriptorSet)
	appendFileDescriptor(m.ParentFile(), fds)
	return fds
}
func appendFileDescriptor(fd protoreflect.FileDescriptor, fds *descriptorpb.FileDescriptorSet) {
	fds.File = append(fds.File, protodesc.ToFileDescriptorProto(fd))
	imps := fd.Imports()
	for i := 0; i < imps.Len(); i++ {
		appendFileDescriptor(imps.Get(i).FileDescriptor, fds)
	}
}
