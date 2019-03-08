// Copyright © 2017 Microsoft <wastore@microsoft.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ste

import (
	"time"

	"github.com/Azure/azure-storage-file-go/azfile"
)

// Source info provider for Azure blob
type fileSourceInfoProvider struct {
	defaultRemoteSourceInfoProvider
}

func newFileSourceInfoProvider(jptm IJobPartTransferMgr) (ISourceInfoProvider, error) {
	b, err := newDefaultRemoteSourceInfoProvider(jptm)
	if err != nil {
		return nil, err
	}

	base, _ := b.(*defaultRemoteSourceInfoProvider)

	return &fileSourceInfoProvider{defaultRemoteSourceInfoProvider: *base}, nil
}

func (p *fileSourceInfoProvider) GetLastModifiedTime() (time.Time, error) {
	presignedURL, err := p.PreSignedSourceURL()
	if err != nil {
		return time.Time{}, err
	}

	fileURL := azfile.NewFileURL(*presignedURL, p.jptm.SourceProviderPipeline())
	properties, err := fileURL.GetProperties(p.jptm.Context())
	if err != nil {
		return time.Time{}, err
	}

	return properties.LastModified(), nil
}
