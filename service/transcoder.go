package main

import (
	"context"
	"fmt"
	"github.com/meschbach/xp-go-dyanmic-transcode/ipc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type Transcoder struct {
	ipc.UnimplementedTranscodeServer
}

func (t *Transcoder) ToJSON(ctx context.Context, in *ipc.JsonIn) (*ipc.JsonOut, error) {
	reg, err := protodesc.NewFiles(in.Format)
	if err != nil {
		panic(err)
	}
	d, err := reg.FindDescriptorByName(protoreflect.FullName(in.What.TypeUrl))
	if err != nil {
		fmt.Printf("Failed to find descriptor for %s\n", in.What.TypeUrl)
		return nil, err
	}
	//todo: d is not guaranteed to actually be a MessageDescriptor
	dyn := dynamicpb.NewMessage(d.(protoreflect.MessageDescriptor))
	if err := in.What.UnmarshalTo(dyn); err != nil {
		panic(err)
	}
	out := protojson.Format(dyn)
	return &ipc.JsonOut{Json: out}, nil
}
