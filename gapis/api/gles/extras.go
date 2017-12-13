// Copyright (C) 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gles

import (
	"context"

	"github.com/google/gapid/core/data/deep"
	"github.com/google/gapid/core/data/id"
	"github.com/google/gapid/core/data/protoconv"
	"github.com/google/gapid/gapis/api"
	"github.com/google/gapid/gapis/api/gles/gles_pb"
)

// ErrorState is a command extra used to describe the GLES error state after
// the command has been executed. It is optional - we use it only for testing.
type ErrorState struct {
	TraceDriversGlError GLenum
	InterceptorsGlError GLenum
}

// EGLImageData is an extra used to store snapshot of external image source.
type EGLImageData struct {
	ID     id.ID
	Size   uint64
	Width  GLsizei
	Height GLsizei
	Format GLenum
	Type   GLenum
}

func init() {
	protoconv.Register(
		func(ctx context.Context, o *ErrorState) (*gles_pb.ErrorState, error) {
			return &gles_pb.ErrorState{
				TraceDriversGlError: uint32(o.TraceDriversGlError),
				InterceptorsGlError: uint32(o.InterceptorsGlError),
			}, nil
		}, func(ctx context.Context, p *gles_pb.ErrorState) (*ErrorState, error) {
			return &ErrorState{
				TraceDriversGlError: GLenum(p.TraceDriversGlError),
				InterceptorsGlError: GLenum(p.InterceptorsGlError),
			}, nil
		},
	)
	protoconv.Register(
		func(ctx context.Context, o *EGLImageData) (*gles_pb.EGLImageData, error) {
			resIndex, err := id.GetRemapper(ctx).RemapID(ctx, o.ID)
			if err != nil {
				return nil, err
			}
			return &gles_pb.EGLImageData{
				ResIndex: resIndex,
				Size:     int32(o.Size),
				Width:    int32(o.Width),
				Height:   int32(o.Height),
				Format:   int32(o.Format),
				Type:     int32(o.Type),
			}, nil
		}, func(ctx context.Context, p *gles_pb.EGLImageData) (*EGLImageData, error) {
			id, err := id.GetRemapper(ctx).RemapIndex(ctx, p.ResIndex)
			if err != nil {
				return nil, err
			}
			return &EGLImageData{
				ID:     id,
				Size:   uint64(p.Size),
				Width:  GLsizei(p.Width),
				Height: GLsizei(p.Height),
				Format: GLenum(p.Format),
				Type:   GLenum(p.Type),
			}, nil
		},
	)
}

// FindCompileShaderExtra searches for the CompileShaderExtra in the extras,
// returning the CompileShaderExtra if found, otherwise nil.
func FindCompileShaderExtra(extras *api.CmdExtras, forShader *Shader) *CompileShaderExtra {
	for _, e := range extras.All() {
		if pi, ok := e.(*CompileShaderExtra); ok {
			// There can be several of those extras - pick the right one.
			if pi.ID == forShader.ID {
				clone, err := deep.Clone(pi)
				if err != nil {
					panic(err)
				}
				return clone.(*CompileShaderExtra)
			}
		}
	}
	return nil
}

// FindLinkProgramExtra searches for the LinkProgramExtra in the extras,
// returning the LinkProgramExtra if found, otherwise nil.
func FindLinkProgramExtra(extras *api.CmdExtras) *LinkProgramExtra {
	for _, e := range extras.All() {
		if pi, ok := e.(*LinkProgramExtra); ok {
			clone, err := deep.Clone(pi)
			if err != nil {
				panic(err)
			}
			return clone.(*LinkProgramExtra)
		}
	}
	return nil
}

// FindValidateProgramExtra searches for the ValidateProgramExtra in the extras,
// returning the ValidateProgramExtra if found, otherwise nil.
func FindValidateProgramExtra(extras *api.CmdExtras) *ValidateProgramExtra {
	for _, e := range extras.All() {
		if pi, ok := e.(*ValidateProgramExtra); ok {
			clone, err := deep.Clone(pi)
			if err != nil {
				panic(err)
			}
			return clone.(*ValidateProgramExtra)
		}
	}
	return nil
}

// FindValidateProgramPipelineExtra searches for the ValidateProgramPipelineExtra in the extras,
// returning the ValidateProgramPipelineExtra if found, otherwise nil.
func FindValidateProgramPipelineExtra(extras *api.CmdExtras) *ValidateProgramPipelineExtra {
	for _, e := range extras.All() {
		if pi, ok := e.(*ValidateProgramPipelineExtra); ok {
			clone, err := deep.Clone(pi)
			if err != nil {
				panic(err)
			}
			return clone.(*ValidateProgramPipelineExtra)
		}
	}
	return nil
}

// FindErrorState searches for the ErrorState in the extras, returning the
// ErrorState if found, otherwise nil.
func FindErrorState(extras *api.CmdExtras) *ErrorState {
	for _, e := range extras.All() {
		if pi, ok := e.(*ErrorState); ok {
			return pi
		}
	}
	return nil
}

// FindEGLImageData searches for the EGLImageData in the extras, returning the
// EGLImageData if found, otherwise nil.
func FindEGLImageData(extras *api.CmdExtras) *EGLImageData {
	for _, e := range extras.All() {
		if res, ok := e.(*EGLImageData); ok {
			clone, err := deep.Clone(res)
			if err != nil {
				panic(err)
			}
			return clone.(*EGLImageData)
		}
	}
	return nil
}

// FindStaticContextState searches for the StaticContextState in the extras,
// returning the StaticContextState if found, otherwise nil.
func FindStaticContextState(extras *api.CmdExtras) *StaticContextState {
	for _, e := range extras.All() {
		if cs, ok := e.(*StaticContextState); ok {
			clone, err := deep.Clone(cs)
			if err != nil {
				panic(err)
			}
			return clone.(*StaticContextState)
		}
	}
	return nil
}

// FindDynamicContextState searches for the DynamicContextState in the extras,
// returning the DynamicContextState if found, otherwise nil.
func FindDynamicContextState(extras *api.CmdExtras) *DynamicContextState {
	for _, e := range extras.All() {
		if cs, ok := e.(*DynamicContextState); ok {
			clone, err := deep.Clone(cs)
			if err != nil {
				panic(err)
			}
			return clone.(*DynamicContextState)
		}
	}
	return nil
}

// FindAndroidNativeBufferExtra searches for the AndroidNativeBufferExtra in the extras,
// returning the AndroidNativeBufferExtra if found, otherwise nil.
func FindAndroidNativeBufferExtra(extras *api.CmdExtras) *AndroidNativeBufferExtra {
	for _, e := range extras.All() {
		if di, ok := e.(*AndroidNativeBufferExtra); ok {
			clone, err := deep.Clone(di)
			if err != nil {
				panic(err)
			}
			return clone.(*AndroidNativeBufferExtra)
		}
	}
	return nil
}